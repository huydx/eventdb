// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package ident

import (
	"github.com/m3db/m3x/pool"
)

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

// tagArrayPool provides a pool for tag slices.
type tagArrayPool interface {
	// Init initializes the array pool, it needs to be called
	// before Get/Put use.
	Init()

	// Get returns the a slice from the pool.
	Get() []Tag

	// Put returns the provided slice to the pool.
	Put(elems []Tag)
}

type tagFinalizeFn func([]Tag) []Tag

type tagArrayPoolOpts struct {
	Options     pool.ObjectPoolOptions
	Capacity    int
	MaxCapacity int
	FinalizeFn  tagFinalizeFn
}

type tagArrPool struct {
	opts tagArrayPoolOpts
	pool pool.ObjectPool
}

func newTagArrayPool(opts tagArrayPoolOpts) tagArrayPool {
	if opts.FinalizeFn == nil {
		opts.FinalizeFn = defaultTagFinalizerFn
	}
	p := pool.NewObjectPool(opts.Options)
	return &tagArrPool{opts, p}
}

func (p *tagArrPool) Init() {
	p.pool.Init(func() interface{} {
		return make([]Tag, 0, p.opts.Capacity)
	})
}

func (p *tagArrPool) Get() []Tag {
	return p.pool.Get().([]Tag)
}

func (p *tagArrPool) Put(arr []Tag) {
	arr = p.opts.FinalizeFn(arr)
	if max := p.opts.MaxCapacity; max > 0 && cap(arr) > max {
		return
	}
	p.pool.Put(arr)
}

func defaultTagFinalizerFn(elems []Tag) []Tag {
	var empty Tag
	for i := range elems {
		elems[i] = empty
	}
	elems = elems[:0]
	return elems
}

type tagArr []Tag

func (elems tagArr) grow(n int) []Tag {
	if cap(elems) < n {
		elems = make([]Tag, n)
	}
	elems = elems[:n]
	// following compiler optimized memcpy impl
	// https://github.com/golang/go/wiki/CompilerOptimizations#optimized-memclr
	var empty Tag
	for i := range elems {
		elems[i] = empty
	}
	return elems
}
