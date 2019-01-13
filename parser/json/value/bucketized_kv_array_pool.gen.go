// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package value

import (
	"fmt"

	"sort"

	"github.com/uber-go/tally"
)

// KVArrayBucket specifies a bucket.
type KVArrayBucket struct {
	// Capacity is the size of each element in the bucket.
	Capacity int

	// Count is the number of fixed elements in the bucket.
	Count int

	// Options is an optional override to specify options to use for a bucket,
	// specify nil to use the options specified to the bucketized pool
	// constructor for this bucket.
	Options *KVArrayPoolOptions
}

// kvArrayBucketByCapacity is a sortable collection of pool buckets.
type kvArrayBucketByCapacity []KVArrayBucket

func (x kvArrayBucketByCapacity) Len() int {
	return len(x)
}

func (x kvArrayBucketByCapacity) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x kvArrayBucketByCapacity) Less(i, j int) bool {
	return x[i].Capacity < x[j].Capacity
}

type kvArrayBucketPool struct {
	capacity int
	pool     *KVArrayPool
}

// BucketizedKVArrayPool is a bucketized value pool.
type BucketizedKVArrayPool struct {
	sizesAsc          []KVArrayBucket
	buckets           []kvArrayBucketPool
	maxBucketCapacity int
	opts              *KVArrayPoolOptions
	alloc             func(capacity int) KVArray
	maxAlloc          tally.Counter
}

// NewBucketizedKVArrayPool creates a bucketized object pool.
func NewBucketizedKVArrayPool(sizes []KVArrayBucket, opts *KVArrayPoolOptions) *BucketizedKVArrayPool {
	if opts == nil {
		opts = NewKVArrayPoolOptions()
	}

	sizesAsc := make([]KVArrayBucket, len(sizes))
	copy(sizesAsc, sizes)
	sort.Sort(kvArrayBucketByCapacity(sizesAsc))

	var maxBucketCapacity int
	if len(sizesAsc) != 0 {
		maxBucketCapacity = sizesAsc[len(sizesAsc)-1].Capacity
	}

	return &BucketizedKVArrayPool{
		opts:              opts,
		sizesAsc:          sizesAsc,
		maxBucketCapacity: maxBucketCapacity,
		maxAlloc:          opts.MetricsScope().Counter("alloc-max"),
	}
}

// Init initializes the bucketized pool.
func (p *BucketizedKVArrayPool) Init(alloc func(capacity int) KVArray) {
	buckets := make([]kvArrayBucketPool, len(p.sizesAsc))
	for i := range p.sizesAsc {
		size := p.sizesAsc[i].Count
		capacity := p.sizesAsc[i].Capacity

		opts := p.opts
		if perBucketOpts := p.sizesAsc[i].Options; perBucketOpts != nil {
			opts = perBucketOpts
		}

		opts = opts.SetSize(size)
		scope := opts.MetricsScope()
		if scope != nil {
			opts = opts.SetMetricsScope(scope.Tagged(map[string]string{
				"bucket-capacity": fmt.Sprintf("%d", capacity),
			}))
		}

		buckets[i].capacity = capacity
		buckets[i].pool = NewKVArrayPool(opts)
		buckets[i].pool.Init(func() KVArray {
			return alloc(capacity)
		})
	}
	p.buckets = buckets
	p.alloc = alloc
}

// Get gets a value from the pool.
func (p *BucketizedKVArrayPool) Get(capacity int) KVArray {
	if capacity > p.maxBucketCapacity {
		p.maxAlloc.Inc(1)
		return p.alloc(capacity)
	}
	for i := range p.buckets {
		if p.buckets[i].capacity >= capacity {
			return p.buckets[i].pool.Get()
		}
	}
	return p.alloc(capacity)
}

// Put puts a value to the pool.
func (p *BucketizedKVArrayPool) Put(v KVArray, capacity int) {
	if capacity > p.maxBucketCapacity {
		return
	}

	for i := len(p.buckets) - 1; i >= 0; i-- {
		if capacity >= p.buckets[i].capacity {
			p.buckets[i].pool.Put(v)
			return
		}
	}
}
