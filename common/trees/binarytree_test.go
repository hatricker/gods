package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	assert.Equal([]int32{1, 3, 2}, Marshal(root))

	root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 2}
	assert.Equal([]int32{1, minInt32, 3, 2}, Marshal(root))

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	assert.Equal([]int32{1, 3, minInt32, minInt32, 2}, Marshal(root))

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	assert.Equal([]int32{1, 3, 2, 4, minInt32, minInt32, 5}, Marshal(root))

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	assert.Equal([]int32{1, 3, 2, minInt32, 4, minInt32, 5}, Marshal(root))
}

func TestUnmarshal(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input, result []int32
	}{
		{
			[]int32{1, 3, 2, 4, minInt32, minInt32, 5},
			[]int32{1, 3, 2, 4, minInt32, minInt32, 5},
		},
		{
			[]int32{1, 3, minInt32, minInt32, 2},
			[]int32{1, 3, minInt32, minInt32, 2},
		},
		{
			[]int32{1, 3, 2, minInt32, minInt32},
			[]int32{1, 3, 2},
		},
	}

	for _, tt := range tests {
		result := Marshal(Unmarshal(tt.input))
		assert.Equal(tt.result, result)
	}
}
