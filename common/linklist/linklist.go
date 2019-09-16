package linklist

//ListNode defines node in single linklist
type ListNode struct {
	Val  int32
	Next *ListNode
}

//Unmarshal builds a single linklist from input integer array
func Unmarshal(values []int32) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	curr := head
	for _, val := range values[1:] {
		node := &ListNode{Val: val}
		curr.Next = node
		curr = node
	}
	return head
}

//Marshal serializes linklist into integer array
func Marshal(node *ListNode) []int32 {
	if node == nil {
		return nil
	}
	result := make([]int32, 0, 64)
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
