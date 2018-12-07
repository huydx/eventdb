// Copyright (c) 2018 Uber Technologies, Inc.
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
	"github.com/uber-go/tally"
)

// Int64ArrayPoolWatermarkConfiguration contains watermark configuration for pools.
type Int64ArrayPoolWatermarkConfiguration struct {
	// The low watermark to start refilling the pool, if zero none.
	RefillLowWatermark float64 `yaml:"low" validate:"min=0.0,max=1.0"`

	// The high watermark to stop refilling the pool, if zero none.
	RefillHighWatermark float64 `yaml:"high" validate:"min=0.0,max=1.0"`
}

// Int64ArrayPoolConfiguration contains pool configuration.
type Int64ArrayPoolConfiguration struct {
	// The size of the pool.
	Size *int `yaml:"size"`

	// The watermark configuration.
	Watermark Int64ArrayPoolWatermarkConfiguration `yaml:"watermark"`
}

// NewPoolOptions creates a new set of pool options.
func (c *Int64ArrayPoolConfiguration) NewPoolOptions(
	scope tally.Scope,
) *Int64ArrayPoolOptions {
	opts := NewInt64ArrayPoolOptions().
		SetMetricsScope(scope).
		SetRefillLowWatermark(c.Watermark.RefillLowWatermark).
		SetRefillHighWatermark(c.Watermark.RefillHighWatermark)
	if c.Size != nil {
		opts = opts.SetSize(*c.Size)
	}
	return opts
}
