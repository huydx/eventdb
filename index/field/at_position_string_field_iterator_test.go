package field

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/xichen2020/eventdb/index"
	"github.com/xichen2020/eventdb/values/iterator"

	"github.com/golang/mock/gomock"
)

func TestNewAtPositionStringFieldIteratorForwardOnly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	docPosIt := index.NewMockDocIDPositionIterator(ctrl)
	gomock.InOrder(
		docPosIt.EXPECT().Next().Return(true),
		docPosIt.EXPECT().Position().Return(0),
		docPosIt.EXPECT().DocID().Return(int32(12)),
		docPosIt.EXPECT().Next().Return(true),
		docPosIt.EXPECT().Position().Return(2),
		docPosIt.EXPECT().DocID().Return(int32(23)),
		docPosIt.EXPECT().Next().Return(true),
		docPosIt.EXPECT().Position().Return(4),
		docPosIt.EXPECT().DocID().Return(int32(45)),
		docPosIt.EXPECT().Next().Return(false),
		docPosIt.EXPECT().Close(),
	)

	valsIt := iterator.NewMockForwardStringIterator(ctrl)
	gomock.InOrder(
		valsIt.EXPECT().Next().Return(true),
		valsIt.EXPECT().Current().Return("a"),
		valsIt.EXPECT().Next().Return(true),
		valsIt.EXPECT().Next().Return(true),
		valsIt.EXPECT().Current().Return("c"),
		valsIt.EXPECT().Next().Return(true),
		valsIt.EXPECT().Next().Return(true),
		valsIt.EXPECT().Current().Return("e"),
		valsIt.EXPECT().Close(),
	)

	var (
		expectedDocIDs = []int32{12, 23, 45}
		expectedValues = []string{"a", "c", "e"}
		actualDocIDs   []int32
		actualValues   []string
	)
	it := newAtPositionStringFieldIterator(docPosIt, valsIt)
	defer it.Close()

	for it.Next() {
		actualDocIDs = append(actualDocIDs, it.DocID())
		actualValues = append(actualValues, it.Value())
	}
	require.NoError(t, it.Err())
	require.Equal(t, expectedDocIDs, actualDocIDs)
	require.Equal(t, expectedValues, actualValues)
}

func TestNewAtPositionStringFieldIteratorSeekable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	docPosIt := index.NewMockDocIDPositionIterator(ctrl)
	gomock.InOrder(
		docPosIt.EXPECT().Next().Return(true),
		docPosIt.EXPECT().Position().Return(0),
		docPosIt.EXPECT().DocID().Return(int32(12)),
		docPosIt.EXPECT().Next().Return(true),
		docPosIt.EXPECT().Position().Return(2),
		docPosIt.EXPECT().DocID().Return(int32(23)),
		docPosIt.EXPECT().Next().Return(true),
		docPosIt.EXPECT().Position().Return(4),
		docPosIt.EXPECT().DocID().Return(int32(45)),
		docPosIt.EXPECT().Next().Return(false),
		docPosIt.EXPECT().Close(),
	)

	valsIt := iterator.NewMockSeekableStringIterator(ctrl)
	gomock.InOrder(
		valsIt.EXPECT().Next().Return(true),
		valsIt.EXPECT().SeekForward(0).Return(nil),
		valsIt.EXPECT().Current().Return("a"),
		valsIt.EXPECT().SeekForward(2).Return(nil),
		valsIt.EXPECT().Current().Return("c"),
		valsIt.EXPECT().SeekForward(2).Return(nil),
		valsIt.EXPECT().Current().Return("e"),
		valsIt.EXPECT().Close(),
	)

	var (
		expectedDocIDs = []int32{12, 23, 45}
		expectedValues = []string{"a", "c", "e"}
		actualDocIDs   []int32
		actualValues   []string
	)
	it := newAtPositionStringFieldIterator(docPosIt, valsIt)
	defer it.Close()

	for it.Next() {
		actualDocIDs = append(actualDocIDs, it.DocID())
		actualValues = append(actualValues, it.Value())
	}
	require.NoError(t, it.Err())
	require.Equal(t, expectedDocIDs, actualDocIDs)
	require.Equal(t, expectedValues, actualValues)
}
