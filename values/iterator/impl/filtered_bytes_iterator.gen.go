// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package impl

import "github.com/xichen2020/eventdb/values/iterator"
import "github.com/xichen2020/eventdb/filter"

// FilteredBytesIterator is a position iterator that outputs the positions
// of values in the value sequence matching a given filter. The position starts at 0.
type FilteredBytesIterator struct {
	vit iterator.ForwardBytesIterator
	f   filter.BytesFilter

	done    bool
	currPos int
	err     error
}

// NewFilteredBytesIterator creates a new filtering iterator.
func NewFilteredBytesIterator(
	vit iterator.ForwardBytesIterator,
	f filter.BytesFilter,
) *FilteredBytesIterator {
	return &FilteredBytesIterator{
		vit:     vit,
		f:       f,
		currPos: -1,
	}
}

// Next returns true if there are more values to be iterated over.
func (it *FilteredBytesIterator) Next() bool {
	if it.done || it.err != nil {
		return false
	}
	for it.vit.Next() {
		it.currPos++
		if it.f.Match(it.vit.Current()) {
			return true
		}
	}
	it.done = true
	it.err = it.vit.Err()
	return false
}

// Position returns the current position.
func (it *FilteredBytesIterator) Position() int { return it.currPos }

// Err returns any errors encountered during iteration.
func (it *FilteredBytesIterator) Err() error { return it.err }

// Close closes the iterator.
func (it *FilteredBytesIterator) Close() {
	it.vit.Close()
	it.vit = nil
	it.f = nil
}