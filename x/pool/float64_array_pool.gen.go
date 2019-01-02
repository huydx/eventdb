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
	"errors"

	"math"

	"sync/atomic"

	"github.com/uber-go/tally"
)

// Float64ArrayPoolOptions provide a set of options for the value pool.
type Float64ArrayPoolOptions struct {
	scope               tally.Scope
	size                int
	refillLowWatermark  float64
	refillHighWatermark float64
}

// NewFloat64ArrayPoolOptions create a new set of value pool options.
func NewFloat64ArrayPoolOptions() *Float64ArrayPoolOptions {
	return &Float64ArrayPoolOptions{
		scope: tally.NoopScope,
		size:  4096,
	}
}

// SetMetricsScope sets the metrics scope.
func (o *Float64ArrayPoolOptions) SetMetricsScope(v tally.Scope) *Float64ArrayPoolOptions {
	opts := *o
	opts.scope = v
	return &opts
}

// MetricsScope returns the metrics scope.
func (o *Float64ArrayPoolOptions) MetricsScope() tally.Scope { return o.scope }

// SetSize sets the pool size.
func (o *Float64ArrayPoolOptions) SetSize(v int) *Float64ArrayPoolOptions {
	opts := *o
	opts.size = v
	return &opts
}

// Size returns pool size.
func (o *Float64ArrayPoolOptions) Size() int { return o.size }

// SetRefillLowWatermark sets the low watermark for refilling the pool.
func (o *Float64ArrayPoolOptions) SetRefillLowWatermark(v float64) *Float64ArrayPoolOptions {
	opts := *o
	opts.refillLowWatermark = v
	return &opts
}

// RefillLowWatermark returns the low watermark for refilling the pool.
func (o *Float64ArrayPoolOptions) RefillLowWatermark() float64 { return o.refillLowWatermark }

// SetRefillHighWatermark sets the high watermark for refilling the pool.
func (o *Float64ArrayPoolOptions) SetRefillHighWatermark(v float64) *Float64ArrayPoolOptions {
	opts := *o
	opts.refillHighWatermark = v
	return &opts
}

// RefillHighWatermark returns the high watermark for stop refilling the pool.
func (o *Float64ArrayPoolOptions) RefillHighWatermark() float64 { return o.refillHighWatermark }

type float64ArrayPoolMetrics struct {
	free       tally.Gauge
	total      tally.Gauge
	getOnEmpty tally.Counter
	putOnFull  tally.Counter
}

func newfloat64ArrayPoolMetrics(m tally.Scope) float64ArrayPoolMetrics {
	return float64ArrayPoolMetrics{
		free:       m.Gauge("free"),
		total:      m.Gauge("total"),
		getOnEmpty: m.Counter("get-on-empty"),
		putOnFull:  m.Counter("put-on-full"),
	}
}

// Float64ArrayPool is a value pool.
type Float64ArrayPool struct {
	values              chan []float64
	alloc               func() []float64
	size                int
	refillLowWatermark  int
	refillHighWatermark int
	filling             int32
	initialized         int32
	dice                int32
	metrics             float64ArrayPoolMetrics
}

// NewFloat64ArrayPool creates a new pool.
func NewFloat64ArrayPool(opts *Float64ArrayPoolOptions) *Float64ArrayPool {
	if opts == nil {
		opts = NewFloat64ArrayPoolOptions()
	}

	p := &Float64ArrayPool{
		values: make(chan []float64, opts.Size()),
		size:   opts.Size(),
		refillLowWatermark: int(math.Ceil(
			opts.RefillLowWatermark() * float64(opts.Size()))),
		refillHighWatermark: int(math.Ceil(
			opts.RefillHighWatermark() * float64(opts.Size()))),
		metrics: newfloat64ArrayPoolMetrics(opts.MetricsScope()),
	}

	p.setGauges()

	return p
}

// Init initializes the pool.
func (p *Float64ArrayPool) Init(alloc func() []float64) {
	if !atomic.CompareAndSwapInt32(&p.initialized, 0, 1) {
		panic(errors.New("pool is already initialized"))
	}

	p.alloc = alloc

	for i := 0; i < cap(p.values); i++ {
		p.values <- p.alloc()
	}

	p.setGauges()
}

// Get gets a value from the pool.
func (p *Float64ArrayPool) Get() []float64 {
	if atomic.LoadInt32(&p.initialized) != 1 {
		panic(errors.New("get before pool is initialized"))
	}

	var v []float64
	select {
	case v = <-p.values:
	default:
		v = p.alloc()
		p.metrics.getOnEmpty.Inc(1)
	}

	p.trySetGauges()

	if p.refillLowWatermark > 0 && len(p.values) <= p.refillLowWatermark {
		p.tryFill()
	}

	return v
}

// Put returns a value to pool.
func (p *Float64ArrayPool) Put(v []float64) {
	if atomic.LoadInt32(&p.initialized) != 1 {
		panic(errors.New("put before pool is initialized"))
	}

	select {
	case p.values <- v:
	default:
		p.metrics.putOnFull.Inc(1)
	}

	p.trySetGauges()
}

func (p *Float64ArrayPool) trySetGauges() {
	if atomic.AddInt32(&p.dice, 1)%100 == 0 {
		p.setGauges()
	}
}

func (p *Float64ArrayPool) setGauges() {
	p.metrics.free.Update(float64(len(p.values)))
	p.metrics.total.Update(float64(p.size))
}

func (p *Float64ArrayPool) tryFill() {
	if !atomic.CompareAndSwapInt32(&p.filling, 0, 1) {
		return
	}

	go func() {
		defer atomic.StoreInt32(&p.filling, 0)

		for len(p.values) < p.refillHighWatermark {
			select {
			case p.values <- p.alloc():
			default:
				return
			}
		}
	}()
}
