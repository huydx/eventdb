// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package impl

import (
	"fmt"

	"github.com/xichen2020/eventdb/document/field"

	"github.com/xichen2020/eventdb/filter"

	"github.com/xichen2020/eventdb/values/iterator"

	iterimpl "github.com/xichen2020/eventdb/values/iterator/impl"
)

// defaultFilteredArrayBasedBoolValueIterator creates a default bool value iterator.
func defaultFilteredArrayBasedBoolValueIterator(
	values *ArrayBasedBoolValues,
	op filter.Op,
	filterValue *field.ValueUnion,
) (iterator.PositionIterator, error) {
	flt, err := op.BoolFilter(filterValue)
	if err != nil {
		return nil, fmt.Errorf("invalid bool filter op %v with filter value %v", op, filterValue)
	}
	valuesIt, err := values.Iter()
	if err != nil {
		return nil, err
	}
	return iterimpl.NewFilteredBoolIterator(valuesIt, flt), nil
}

// defaultFilteredArrayBasedIntValueIterator creates a default int value iterator.
func defaultFilteredArrayBasedIntValueIterator(
	values *ArrayBasedIntValues,
	op filter.Op,
	filterValue *field.ValueUnion,
) (iterator.PositionIterator, error) {
	flt, err := op.IntFilter(filterValue)
	if err != nil {
		return nil, fmt.Errorf("invalid int filter op %v with filter value %v", op, filterValue)
	}
	valuesIt, err := values.Iter()
	if err != nil {
		return nil, err
	}
	return iterimpl.NewFilteredIntIterator(valuesIt, flt), nil
}

// defaultFilteredArrayBasedDoubleValueIterator creates a default double value iterator.
func defaultFilteredArrayBasedDoubleValueIterator(
	values *ArrayBasedDoubleValues,
	op filter.Op,
	filterValue *field.ValueUnion,
) (iterator.PositionIterator, error) {
	flt, err := op.DoubleFilter(filterValue)
	if err != nil {
		return nil, fmt.Errorf("invalid double filter op %v with filter value %v", op, filterValue)
	}
	valuesIt, err := values.Iter()
	if err != nil {
		return nil, err
	}
	return iterimpl.NewFilteredDoubleIterator(valuesIt, flt), nil
}

// defaultFilteredArrayBasedBytesValueIterator creates a default bytes value iterator.
func defaultFilteredArrayBasedBytesValueIterator(
	values *ArrayBasedBytesValues,
	op filter.Op,
	filterValue *field.ValueUnion,
) (iterator.PositionIterator, error) {
	flt, err := op.BytesFilter(filterValue)
	if err != nil {
		return nil, fmt.Errorf("invalid bytes filter op %v with filter value %v", op, filterValue)
	}
	valuesIt, err := values.Iter()
	if err != nil {
		return nil, err
	}
	return iterimpl.NewFilteredBytesIterator(valuesIt, flt), nil
}

// defaultFilteredArrayBasedTimeValueIterator creates a default time value iterator.
func defaultFilteredArrayBasedTimeValueIterator(
	values *ArrayBasedTimeValues,
	op filter.Op,
	filterValue *field.ValueUnion,
) (iterator.PositionIterator, error) {
	flt, err := op.TimeFilter(filterValue)
	if err != nil {
		return nil, fmt.Errorf("invalid time filter op %v with filter value %v", op, filterValue)
	}
	valuesIt, err := values.Iter()
	if err != nil {
		return nil, err
	}
	return iterimpl.NewFilteredTimeIterator(valuesIt, flt), nil
}
