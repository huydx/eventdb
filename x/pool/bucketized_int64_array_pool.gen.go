// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package pool

import (
	"fmt"

	"sort"

	"github.com/uber-go/tally"
)

// Int64ArrayBucket specifies a bucket.
type Int64ArrayBucket struct {
	// Capacity is the size of each element in the bucket.
	Capacity int

	// Count is the number of fixed elements in the bucket.
	Count int

	// Options is an optional override to specify options to use for a bucket,
	// specify nil to use the options specified to the bucketized pool
	// constructor for this bucket.
	Options *Int64ArrayPoolOptions
}

// int64ArrayBucketByCapacity is a sortable collection of pool buckets.
type int64ArrayBucketByCapacity []Int64ArrayBucket

func (x int64ArrayBucketByCapacity) Len() int {
	return len(x)
}

func (x int64ArrayBucketByCapacity) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x int64ArrayBucketByCapacity) Less(i, j int) bool {
	return x[i].Capacity < x[j].Capacity
}

type int64ArrayBucketPool struct {
	capacity int
	pool     *Int64ArrayPool
}

// BucketizedInt64ArrayPool is a bucketized value pool.
type BucketizedInt64ArrayPool struct {
	sizesAsc          []Int64ArrayBucket
	buckets           []int64ArrayBucketPool
	maxBucketCapacity int
	opts              *Int64ArrayPoolOptions
	alloc             func(capacity int) []int64
	maxAlloc          tally.Counter
}

// NewBucketizedInt64ArrayPool creates a bucketized object pool.
func NewBucketizedInt64ArrayPool(sizes []Int64ArrayBucket, opts *Int64ArrayPoolOptions) *BucketizedInt64ArrayPool {
	if opts == nil {
		opts = NewInt64ArrayPoolOptions()
	}

	sizesAsc := make([]Int64ArrayBucket, len(sizes))
	copy(sizesAsc, sizes)
	sort.Sort(int64ArrayBucketByCapacity(sizesAsc))

	var maxBucketCapacity int
	if len(sizesAsc) != 0 {
		maxBucketCapacity = sizesAsc[len(sizesAsc)-1].Capacity
	}

	return &BucketizedInt64ArrayPool{
		opts:              opts,
		sizesAsc:          sizesAsc,
		maxBucketCapacity: maxBucketCapacity,
		maxAlloc:          opts.MetricsScope().Counter("alloc-max"),
	}
}

// Init initializes the bucketized pool.
func (p *BucketizedInt64ArrayPool) Init(alloc func(capacity int) []int64) {
	buckets := make([]int64ArrayBucketPool, len(p.sizesAsc))
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
		buckets[i].pool = NewInt64ArrayPool(opts)
		buckets[i].pool.Init(func() []int64 {
			return alloc(capacity)
		})
	}
	p.buckets = buckets
	p.alloc = alloc
}

// Get gets a value from the pool.
func (p *BucketizedInt64ArrayPool) Get(capacity int) []int64 {
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
func (p *BucketizedInt64ArrayPool) Put(v []int64, capacity int) {
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
