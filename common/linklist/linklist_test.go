package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	assert := assert.New(t)

	values := []int32{}
	assert.Equal((*ListNode)(nil), Unmarshal(values))

	values = []int32{1, 2}
	head := Unmarshal(values)
	assert.Equal(int32(1), head.Val)
	assert.Equal(int32(2), head.Next.Val)
}

func TestMarshal(t *testing.T) {
	assert := assert.New(t)

	head := (*ListNode)(nil)
	assert.Equal(([]int32)(nil), Marshal(head))

	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	assert.Equal([]int32{1, 2}, Marshal(head))
}
