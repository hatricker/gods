package trees

import "github.com/hatricker/gods/common"

const (
	minInt32 = common.MinInt32
)

//TreeNode defines a binary tree node
type TreeNode struct {
	Val         int32
	Left, Right *TreeNode
}

//Unmarshal builds a binary tree from a series of integers
func Unmarshal(values []int32) *TreeNode {
	length := len(values)
	if length == 0 || values[0] == minInt32 {
		return nil
	}
	queue := make([]*TreeNode, 0, length)
	root := &TreeNode{Val: values[0]}
	queue = append(queue, root)
	for index := 1; index < length; index++ {
		node := queue[0]
		if values[index] != minInt32 {
			node.Left = &TreeNode{Val: values[index]}
			queue = append(queue, node.Left)
		}
		index++
		if index == length {
			break
		}

		if values[index] != minInt32 {
			node.Right = &TreeNode{Val: values[index]}
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return root
}

//appendNodeVal appends node value to serialized array including nils ahead of current node
func appendNodeVal(a []int32, nodeVal int32, numNils int) []int32 {
	if numNils > 0 {
		tail := make([]int32, numNils)
		for i := range tail {
			tail[i] = minInt32
		}
		a = append(a, tail...)
	}
	a = append(a, nodeVal)
	return a
}

//Marshal serializes binary tree in BFS order
func Marshal(root *TreeNode) []int32 {
	if root == nil || root.Val == minInt32 {
		return nil
	}
	queue := make([]*TreeNode, 0, 64)
	result := make([]int32, 0, 64)

	queue = append(queue, root)
	result = append(result, root.Val)
	numNils := 0

	for len(queue) > 0 {
		top := queue[0]
		if top.Left != nil {
			queue = append(queue, top.Left)
			result = appendNodeVal(result, top.Left.Val, numNils)
			numNils = 0
		} else {
			numNils++
		}

		if top.Right != nil {
			queue = append(queue, top.Right)
			result = appendNodeVal(result, top.Right.Val, numNils)
			numNils = 0
		} else {
			numNils++
		}
		queue = queue[1:]
	}
	return result
}
