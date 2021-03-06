// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package impl

import "github.com/xichen2020/eventdb/values/iterator"
import "github.com/xichen2020/eventdb/filter"

// FilteredIntIterator is a position iterator that outputs the positions
// of values in the value sequence matching a given filter. The position starts at 0.
type FilteredIntIterator struct {
	vit iterator.ForwardIntIterator
	f   filter.IntFilter

	done    bool
	currPos int
	err     error
}

// NewFilteredIntIterator creates a new filtering iterator.
func NewFilteredIntIterator(
	vit iterator.ForwardIntIterator,
	f filter.IntFilter,
) *FilteredIntIterator {
	return &FilteredIntIterator{
		vit:     vit,
		f:       f,
		currPos: -1,
	}
}

// Next returns true if there are more values to be iterated over.
func (it *FilteredIntIterator) Next() bool {
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
func (it *FilteredIntIterator) Position() int { return it.currPos }

// Err returns any errors encountered during iteration.
func (it *FilteredIntIterator) Err() error { return it.err }

// Close closes the iterator.
func (it *FilteredIntIterator) Close() {
	it.vit.Close()
	it.vit = nil
	it.f = nil
}
