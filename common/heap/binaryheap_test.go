package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapInit(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                      []int32
		maxHeapElems, minHeapElems []int32
	}{
		{
			[]int32{3, 7, 10, 5, 6, 8, 15, 2},
			[]int32{15, 6, 10, 3, 5, 7, 8, 2},
			[]int32{2, 3, 8, 5, 6, 10, 15, 7},
		},
		{
			[]int32{},
			[]int32{},
			[]int32{},
		},
		{
			[]int32{1},
			[]int32{1},
			[]int32{1},
		},
	}

	for _, tt := range tests {
		maxHeap := Init(tt.input, true)
		assert.Equal(tt.maxHeapElems, maxHeap.Elems)

		minHeap := Init(tt.input, false)
		assert.Equal(tt.minHeapElems, minHeap.Elems)
	}
}

func TestHeapInsert(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		currElems                  []int32
		input                      int32
		maxHeapElems, minHeapElems []int32
	}{
		{
			[]int32{3, 7, 10, 5, 6, 8, 15, 2},
			12,
			[]int32{15, 12, 10, 6, 5, 7, 8, 2, 3},
			[]int32{2, 3, 8, 5, 6, 10, 15, 7, 12},
		},
		{
			[]int32{},
			1,
			[]int32{1},
			[]int32{1},
		},
		{
			[]int32{1},
			2,
			[]int32{2, 1},
			[]int32{1, 2},
		},
	}

	for _, tt := range tests {
		maxHeap := Init(tt.currElems, true)
		maxHeap.Insert(tt.input)
		assert.Equal(tt.maxHeapElems, maxHeap.Elems)

		minHeap := Init(tt.currElems, false)
		minHeap.Insert(tt.input)
		assert.Equal(tt.minHeapElems, minHeap.Elems)
	}
}

func TestHeapRemove(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		currElems                  []int32
		maxHeapElems, minHeapElems []int32
	}{
		{
			[]int32{3, 7, 10, 5, 6, 8, 15, 2},
			[]int32{10, 6, 8, 3, 5, 7, 2},
			[]int32{3, 5, 8, 7, 6, 10, 15},
		},
		{
			[]int32{3, 7, 10, 5, 6, 8, 15, 2, 12},
			[]int32{12, 6, 10, 3, 5, 7, 8, 2},
			[]int32{3, 5, 8, 7, 6, 10, 15, 12},
		},
		{
			[]int32{},
			[]int32{},
			[]int32{},
		},
		{
			[]int32{1},
			[]int32{},
			[]int32{},
		},
		{
			[]int32{1, 2},
			[]int32{1},
			[]int32{2},
		},
	}

	for _, tt := range tests {
		maxHeap := Init(tt.currElems, true)
		maxHeap.Remove()
		assert.Equal(tt.maxHeapElems, maxHeap.Elems)

		minHeap := Init(tt.currElems, false)
		minHeap.Remove()
		assert.Equal(tt.minHeapElems, minHeap.Elems)
	}
}
