package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
		n := 2
		expected := 5
		removeNthFromEnd(head, n)

		assert.Equal(t, expected, head.Next.Next.Next.Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{Val: 1}
		n := 1

		assert.Nil(t, removeNthFromEnd(head, n))
	})

	t.Run("test case 3", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
		n := 1
		expected := head

		assert.Equal(t, expected, removeNthFromEnd(head, n))
		assert.Nil(t, head.Next)
	})

	t.Run("test case 4", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
		n := 2
		expected := head.Next

		assert.Equal(t, expected, removeNthFromEnd(head, n))
	})
}
