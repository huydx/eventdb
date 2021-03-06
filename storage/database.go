package storage

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/xichen2020/eventdb/document"
	"github.com/xichen2020/eventdb/query"
	"github.com/xichen2020/eventdb/x/hash"
	"github.com/xichen2020/eventdb/x/safe"

	"github.com/m3db/m3x/clock"
	xerrors "github.com/m3db/m3x/errors"
	"github.com/m3db/m3x/instrument"
	"github.com/uber-go/tally"
)

// Database is a database for timestamped events.
type Database interface {
	// Options returns database options.
	Options() *Options

	// Open opens the database for reading and writing events.
	Open() error

	// WriteBatch writes a batch of timestamped documents to a namespace.
	WriteBatch(
		ctx context.Context,
		namespace []byte,
		docs []document.Document,
	) error

	// QueryRaw executes a raw query against database for documents matching
	// certain criteria, with optional filtering, sorting, and limiting applied.
	QueryRaw(
		ctx context.Context,
		q query.ParsedRawQuery,
	) (*query.RawResults, error)

	// QueryGrouped executes a grouped query against database for documents matching
	// certain criteria, with optional filtering, grouping, sorting, and limiting applied.
	QueryGrouped(
		ctx context.Context,
		q query.ParsedGroupedQuery,
	) (*query.GroupedResults, error)

	// QueryTimeBucket executes a time bucket query against database for documents
	// matching criteria, and counts the number of matching documents in each time bucket.
	QueryTimeBucket(
		ctx context.Context,
		q query.ParsedTimeBucketQuery,
	) (*query.TimeBucketResults, error)

	// Close closes the database.
	Close() error
}

type databaseMetrics struct {
	queryRaw        instrument.MethodMetrics
	queryGrouped    instrument.MethodMetrics
	queryTimeBucket instrument.MethodMetrics
	write           instrument.MethodMetrics
	writeBatch      instrument.MethodMetrics
}

func newDatabaseMetrics(
	scope tally.Scope,
	samplingRate float64,
) databaseMetrics {
	return databaseMetrics{
		queryRaw:        instrument.NewMethodMetrics(scope, "query-raw", samplingRate),
		queryGrouped:    instrument.NewMethodMetrics(scope, "query-grouped", samplingRate),
		queryTimeBucket: instrument.NewMethodMetrics(scope, "query-time-bucket", samplingRate),
		write:           instrument.NewMethodMetrics(scope, "write", samplingRate),
		writeBatch:      instrument.NewMethodMetrics(scope, "write-batch", samplingRate),
	}
}

// database provides internal database APIs.
type database interface {
	Database

	// GetOwnedNamespaces returns the namespaces owned by the database.
	GetOwnedNamespaces() ([]databaseNamespace, error)
}

var (
	// errDatabaseNotOpenOrClosed raised when trying to perform an action that requires
	// the databse is open.
	errDatabaseNotOpenOrClosed = errors.New("database is not open")

	// errDatabaseOpenOrClosed raised when trying to perform an action that requires
	// the database is not open.
	errDatabaseOpenOrClosed = errors.New("database is open or closed")
)

type databaseState int

const (
	databaseNotOpen databaseState = iota
	databaseOpen
	databaseClosed
)

type db struct {
	sync.RWMutex

	opts *Options

	state      databaseState
	mediator   databaseMediator
	flushMgr   databaseFlushManager
	namespaces map[hash.Hash]databaseNamespace

	nowFn   clock.NowFn
	metrics databaseMetrics
}

// NewDatabase creates a new database.
func NewDatabase(
	namespaces []NamespaceMetadata,
	opts *Options,
) Database {
	if opts == nil {
		opts = NewOptions()
	}

	// Create Flush manager.
	instrumentOpts := opts.InstrumentOptions()
	scope := instrumentOpts.MetricsScope()
	flushManagerScope := scope.SubScope("flush-mgr")
	flushManager := newFlushManager(opts.SetInstrumentOptions(instrumentOpts.SetMetricsScope(flushManagerScope)))
	opts = opts.setDatabaseFlushManager(flushManager)

	nss := make(map[hash.Hash]databaseNamespace, len(namespaces))
	for _, ns := range namespaces {
		h := hash.BytesHash(ns.ID())
		nss[h] = newDatabaseNamespace(ns, opts)
	}

	samplingRate := instrumentOpts.MetricsSamplingRate()
	d := &db{
		opts:       opts,
		flushMgr:   flushManager,
		namespaces: nss,
		nowFn:      opts.ClockOptions().NowFn(),
		metrics:    newDatabaseMetrics(scope, samplingRate),
	}
	d.mediator = newMediator(d, opts)

	return d
}

func (d *db) Options() *Options { return d.opts }

func (d *db) Open() error {
	d.Lock()
	defer d.Unlock()

	if d.state != databaseNotOpen {
		return errDatabaseOpenOrClosed
	}
	if err := d.flushMgr.Open(); err != nil {
		return err
	}
	if err := d.mediator.Open(); err != nil {
		return err
	}
	d.state = databaseOpen
	return nil
}

func (d *db) WriteBatch(
	ctx context.Context,
	namespace []byte,
	docs []document.Document,
) error {
	callStart := d.nowFn()
	n, err := d.namespaceFor(namespace)
	if err != nil {
		d.metrics.writeBatch.ReportError(d.nowFn().Sub(callStart))
		return err
	}
	err = n.WriteBatch(ctx, docs)
	d.metrics.writeBatch.ReportSuccessOrError(err, d.nowFn().Sub(callStart))
	return err
}

func (d *db) QueryRaw(
	ctx context.Context,
	q query.ParsedRawQuery,
) (*query.RawResults, error) {
	callStart := d.nowFn()
	n, err := d.namespaceFor(safe.ToBytes(q.Namespace))
	if err != nil {
		d.metrics.queryRaw.ReportError(d.nowFn().Sub(callStart))
		return nil, err
	}
	res, err := n.QueryRaw(ctx, q)
	d.metrics.queryRaw.ReportSuccessOrError(err, d.nowFn().Sub(callStart))
	return res, err
}

func (d *db) QueryGrouped(
	ctx context.Context,
	q query.ParsedGroupedQuery,
) (*query.GroupedResults, error) {
	callStart := d.nowFn()
	n, err := d.namespaceFor(safe.ToBytes(q.Namespace))
	if err != nil {
		return nil, err
	}
	res, err := n.QueryGrouped(ctx, q)
	d.metrics.queryGrouped.ReportSuccessOrError(err, d.nowFn().Sub(callStart))
	return res, err
}

func (d *db) QueryTimeBucket(
	ctx context.Context,
	q query.ParsedTimeBucketQuery,
) (*query.TimeBucketResults, error) {
	callStart := d.nowFn()
	n, err := d.namespaceFor(safe.ToBytes(q.Namespace))
	if err != nil {
		return nil, err
	}
	res, err := n.QueryTimeBucket(ctx, q)
	d.metrics.queryTimeBucket.ReportSuccessOrError(err, d.nowFn().Sub(callStart))
	return res, err
}

func (d *db) Close() error {
	d.Lock()
	defer d.Unlock()

	if d.state != databaseOpen {
		return errDatabaseNotOpenOrClosed
	}
	d.state = databaseClosed

	// Close database-level resources.
	var multiErr xerrors.MultiError
	if err := d.mediator.Close(); err != nil {
		multiErr = multiErr.Add(err)
	}
	if err := d.flushMgr.Close(); err != nil {
		multiErr = multiErr.Add(err)
	}

	// Close namespaces.
	for _, ns := range d.ownedNamespacesWithLock() {
		if err := ns.Close(); err != nil {
			multiErr = multiErr.Add(err)
		}
	}

	return multiErr.FinalError()
}

func (d *db) GetOwnedNamespaces() ([]databaseNamespace, error) {
	d.RLock()
	defer d.RUnlock()
	if d.state != databaseOpen {
		return nil, errDatabaseNotOpenOrClosed
	}
	return d.ownedNamespacesWithLock(), nil
}

func (d *db) namespaceFor(namespace []byte) (databaseNamespace, error) {
	h := hash.BytesHash(namespace)
	d.RLock()
	n, exists := d.namespaces[h]
	d.RUnlock()

	if !exists {
		return nil, fmt.Errorf("no such namespace %s", namespace)
	}
	return n, nil
}

// ownedNamespacesWithLock returns the list of owned namespaces within a lock.
// This ensures the internal list of namespaces may not be modified by the caller.
func (d *db) ownedNamespacesWithLock() []databaseNamespace {
	namespaces := make([]databaseNamespace, 0, len(d.namespaces))
	for _, n := range d.namespaces {
		namespaces = append(namespaces, n)
	}
	return namespaces
}
