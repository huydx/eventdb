package index

import (
	"errors"

	"github.com/xichen2020/eventdb/values"
	"github.com/xichen2020/eventdb/x/refcnt"
)

// IntField contains data in documents for which such field are int values.
// TODO(xichen): Potentially support query APIs.
type IntField interface {
	// DocIDSet returns the doc ID set for which the documents have int values.
	DocIDSet() DocIDSet

	// Values return the collection of int values. The values collection remains
	// valid until the field is closed.
	Values() values.IntValues
}

// CloseableIntField is a int field that can be closed.
type CloseableIntField interface {
	IntField

	// ShallowCopy returns a shallow copy of the field sharing access to the
	// underlying resources. As such the resources held will not be released until
	// there are no more references to the field.
	ShallowCopy() CloseableIntField

	// Close closes the field to release the resources held for the collection.
	Close()
}

// intFieldBuilder incrementally builds a int field.
type intFieldBuilder interface {
	// Add adds a int value alongside its document ID.
	Add(docID int32, v int) error

	// Snapshot take a snapshot of the field data accummulated so far.
	Snapshot() CloseableIntField

	// Seal seals and closes the int builder and returns an immutable int field.
	// The resource ownership is transferred from the builder to the immutable
	// collection as a result. Adding more data to the builder after the builder
	// is sealed will result in an error.
	Seal(numTotalDocs int32) CloseableIntField

	// Close closes the builder.
	Close()
}

var (
	errIntFieldBuilderAlreadyClosed = errors.New("int field builder is already closed")
)

type intField struct {
	*refcnt.RefCounter

	docIDSet DocIDSet
	values   values.CloseableIntValues
	closeFn  FieldCloseFn

	closed bool
}

// NewCloseableIntField creates a int field.
func NewCloseableIntField(
	docIDSet DocIDSet,
	values values.CloseableIntValues,
) CloseableIntField {
	return NewCloseableIntFieldWithCloseFn(docIDSet, values, nil)
}

// NewCloseableIntFieldWithCloseFn creates a int field with a close function.
func NewCloseableIntFieldWithCloseFn(
	docIDSet DocIDSet,
	values values.CloseableIntValues,
	closeFn FieldCloseFn,
) CloseableIntField {
	return &intField{
		RefCounter: refcnt.NewRefCounter(),
		docIDSet:   docIDSet,
		values:     values,
		closeFn:    closeFn,
	}
}

func (f *intField) DocIDSet() DocIDSet       { return f.docIDSet }
func (f *intField) Values() values.IntValues { return f.values }

func (f *intField) ShallowCopy() CloseableIntField {
	f.IncRef()
	shallowCopy := *f
	return &shallowCopy
}

func (f *intField) Close() {
	if f.closed {
		return
	}
	f.closed = true
	if f.DecRef() > 0 {
		return
	}
	f.docIDSet = nil
	f.values.Close()
	f.values = nil
	if f.closeFn != nil {
		f.closeFn()
		f.closeFn = nil
	}
}

type builderOfIntField struct {
	dsb docIDSetBuilder
	svb values.IntValuesBuilder

	closed bool
}

func newIntFieldBuilder(
	dsb docIDSetBuilder,
	svb values.IntValuesBuilder,
) *builderOfIntField {
	return &builderOfIntField{dsb: dsb, svb: svb}
}

func (b *builderOfIntField) Add(docID int32, v int) error {
	if b.closed {
		return errIntFieldBuilderAlreadyClosed
	}
	b.dsb.Add(docID)
	return b.svb.Add(v)
}

func (b *builderOfIntField) Snapshot() CloseableIntField {
	docIDSetSnapshot := b.dsb.Snapshot()
	intValuesSnapshot := b.svb.Snapshot()
	return NewCloseableIntField(docIDSetSnapshot, intValuesSnapshot)
}

func (b *builderOfIntField) Seal(numTotalDocs int32) CloseableIntField {
	docIDSet := b.dsb.Seal(numTotalDocs)
	values := b.svb.Seal()
	sealed := NewCloseableIntField(docIDSet, values)

	// Clear and close the builder so it's no longer writable.
	*b = builderOfIntField{}
	b.Close()

	return sealed
}

func (b *builderOfIntField) Close() {
	if b.closed {
		return
	}
	b.closed = true
	b.dsb = nil
	if b.svb != nil {
		b.svb.Close()
		b.svb = nil
	}
}
