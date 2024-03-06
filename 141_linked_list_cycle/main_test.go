package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
		second := head.Next
		last := head.Next.Next.Next
		last.Next = second

		assert.True(t, hasCycle(head))
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
		last := head.Next
		last.Next = head

		assert.True(t, hasCycle(head))
	})

	t.Run("test case 3", func(t *testing.T) {
		head := &ListNode{Val: 1}

		assert.False(t, hasCycle(head))
	})

	t.Run("test case 4", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}

		assert.False(t, hasCycle(head))
	})
}
