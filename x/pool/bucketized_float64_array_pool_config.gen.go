// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package pool

import (
	"github.com/m3db/m3x/instrument"
)

// Float64ArrayBucketConfiguration contains configuration for a pool bucket.
type Float64ArrayBucketConfiguration struct {
	// The count of the items in the bucket.
	Count int `yaml:"count"`

	// The capacity of each item in the bucket.
	Capacity int `yaml:"capacity"`
}

// NewBucket creates a new bucket.
func (c *Float64ArrayBucketConfiguration) NewBucket() Float64ArrayBucket {
	return Float64ArrayBucket{
		Capacity: c.Capacity,
		Count:    c.Count,
	}
}

// BucketizedFloat64ArrayPoolConfiguration contains configuration for bucketized pools.
type BucketizedFloat64ArrayPoolConfiguration struct {
	// The pool bucket configuration.
	Buckets []Float64ArrayBucketConfiguration `yaml:"buckets"`

	// The watermark configuration.
	Watermark Float64ArrayPoolWatermarkConfiguration `yaml:"watermark"`
}

// NewPoolOptions creates a new set of pool options.
func (c *BucketizedFloat64ArrayPoolConfiguration) NewPoolOptions(
	instrumentOptions instrument.Options,
) *Float64ArrayPoolOptions {
	return NewFloat64ArrayPoolOptions().
		SetInstrumentOptions(instrumentOptions).
		SetRefillLowWatermark(c.Watermark.RefillLowWatermark).
		SetRefillHighWatermark(c.Watermark.RefillHighWatermark)
}

// NewBuckets create a new list of buckets.
func (c *BucketizedFloat64ArrayPoolConfiguration) NewBuckets() []Float64ArrayBucket {
	buckets := make([]Float64ArrayBucket, 0, len(c.Buckets))
	for _, bconfig := range c.Buckets {
		bucket := bconfig.NewBucket()
		buckets = append(buckets, bucket)
	}
	return buckets
}
