// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package query

// timeResultGroupHeap is a heap storing a list of values.
// The ordering of such items are determined by `lessThanFn`.
// The smallest item will be at the top of the heap.
type timeResultGroupHeap struct {
	dv         []timeResultGroup
	lessThanFn func(v1, v2 timeResultGroup) bool
}

// newTimeResultGroupHeap creates a new values heap.
func newTimeResultGroupHeap(
	initCapacity int,
	lessThanFn func(v1, v2 timeResultGroup) bool,
) *timeResultGroupHeap {
	return &timeResultGroupHeap{
		dv:         make([]timeResultGroup, 0, initCapacity),
		lessThanFn: lessThanFn,
	}
}

// RawData returns the underlying backing array in no particular order.
func (h timeResultGroupHeap) RawData() []timeResultGroup { return h.dv }

// Min returns the "smallest" heap element according to the `lessThan` function.
func (h timeResultGroupHeap) Min() timeResultGroup { return h.dv[0] }

// Len returns the number of items in the heap.
func (h timeResultGroupHeap) Len() int { return len(h.dv) }

// Cap returns the heap capacity before a reallocation is needed.
func (h timeResultGroupHeap) Cap() int { return cap(h.dv) }

// Less returns true if item `i` is less than item `j`.
func (h timeResultGroupHeap) Less(i, j int) bool {
	return h.lessThanFn(h.dv[i], h.dv[j])
}

// Swap swaps item `i` with item `j`.
func (h timeResultGroupHeap) Swap(i, j int) { h.dv[i], h.dv[j] = h.dv[j], h.dv[i] }

// Reset resets the internal backing array.
func (h *timeResultGroupHeap) Reset() { h.dv = h.dv[:0] }

// Push pushes a value onto the heap.
func (h *timeResultGroupHeap) Push(value timeResultGroup) {
	h.dv = append(h.dv, value)
	h.shiftUp(h.Len() - 1)
}

// Pop pops a value from the heap.
func (h *timeResultGroupHeap) Pop() timeResultGroup {
	var (
		n   = h.Len()
		val = h.dv[0]
	)

	h.dv[0], h.dv[n-1] = h.dv[n-1], h.dv[0]
	h.heapify(0, n-1)
	h.dv = h.dv[0 : n-1]
	return val
}

// SortInPlace sorts the heap in place and returns the sorted data, with the smallest element
// at the end of the returned array. This is done by repeated swapping the smallest element with
// the last element of the current heap and shrinking the heap size.
// NB: The heap becomes invalid after this is called.
func (h *timeResultGroupHeap) SortInPlace() []timeResultGroup {
	numElems := len(h.dv)
	for len(h.dv) > 0 {
		h.Pop()
	}
	res := h.dv[:numElems]
	h.dv = nil
	h.lessThanFn = nil
	return res
}

func (h timeResultGroupHeap) shiftUp(i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || !h.Less(i, parent) {
			break
		}
		h.dv[parent], h.dv[i] = h.dv[i], h.dv[parent]
		i = parent
	}
}

func (h timeResultGroupHeap) heapify(i, n int) {
	for {
		left := i*2 + 1
		right := left + 1
		smallest := i
		if left < n && h.Less(left, smallest) {
			smallest = left
		}
		if right < n && h.Less(right, smallest) {
			smallest = right
		}
		if smallest == i {
			return
		}
		h.dv[i], h.dv[smallest] = h.dv[smallest], h.dv[i]
		i = smallest
	}
}

// topNTimes keeps track of the top n values in a value sequence for the
// order defined by the `lessThanFn`. In particular if `lessThanFn` defines
// an increasing order (returning true if `v1` < `v2`), the collection stores
// the top N largest values, and vice versa.
type topNTimes struct {
	n          int
	lessThanFn func(v1, v2 timeResultGroup) bool
	h          *timeResultGroupHeap
}

// newTopNTimes creates a new top n value collection.
func newTopNTimes(
	n int,
	lessThanFn func(v1, v2 timeResultGroup) bool,
) *topNTimes {
	return &topNTimes{
		n:          n,
		lessThanFn: lessThanFn,
		h:          newTimeResultGroupHeap(n, lessThanFn),
	}
}

// timeAddOptions provide the options for adding a value.
type timeAddOptions struct {
	CopyOnAdd bool
	CopyFn    func(v timeResultGroup) timeResultGroup
	CopyToFn  func(src timeResultGroup, target *timeResultGroup)
}

// Len returns the number of items in the collection.
func (v topNTimes) Len() int { return v.h.Len() }

// Cap returns the collection capacity.
func (v topNTimes) Cap() int { return v.h.Cap() }

// RawData returns the underlying array backing the heap in no particular order.
func (v topNTimes) RawData() []timeResultGroup { return v.h.RawData() }

// Top returns the "smallest" value according to the `lessThan` function.
func (v topNTimes) Top() timeResultGroup { return v.h.Min() }

// Reset resets the internal array backing the heap.
func (v *topNTimes) Reset() { v.h.Reset() }

// Add adds a value to the collection.
func (v *topNTimes) Add(val timeResultGroup, opts timeAddOptions) {
	if v.h.Len() < v.n {
		if opts.CopyOnAdd {
			val = opts.CopyFn(val)
		}
		v.h.Push(val)
		return
	}
	if min := v.h.Min(); !v.lessThanFn(min, val) {
		return
	}
	popped := v.h.Pop()
	if !opts.CopyOnAdd {
		v.h.Push(val)
		return
	}
	// Reuse popped item from the heap.
	opts.CopyToFn(val, &popped)
	v.h.Push(popped)
}

// SortInPlace sorts the backing heap in place and returns the sorted data.
// NB: The value collection becomes invalid after this is called.
func (v *topNTimes) SortInPlace() []timeResultGroup {
	res := v.h.SortInPlace()
	v.h = nil
	v.lessThanFn = nil
	return res
}
