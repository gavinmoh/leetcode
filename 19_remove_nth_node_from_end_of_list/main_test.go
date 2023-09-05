package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// 1->2->3->4->5, and n = 2.
	case1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	case1Result := RemoveNthFromEnd(case1, 2)
	assert.Equal(t, 5, case1Result.Next.Next.Next.Val)

	case2 := &ListNode{Val: 1}
	case2Result := RemoveNthFromEnd(case2, 1)
	assert.Nil(t, case2Result)

	case3 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	case3Result := RemoveNthFromEnd(case3, 1)
	assert.Equal(t, 1, case3Result.Val)
}
