package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		node5 := &ListNode{Val: 5}
		node4 := &ListNode{Val: 4, Next: node5}
		node3 := &ListNode{Val: 3, Next: node4}
		node2 := &ListNode{Val: 2, Next: node3}
		node1 := &ListNode{Val: 1, Next: node2}
		k := 2
		newHead := reverseKGroup(node1, k)
		assert.Equal(t, node2, newHead)
		assert.Equal(t, node1, newHead.Next)
		assert.Equal(t, node4, newHead.Next.Next)
		assert.Equal(t, node3, newHead.Next.Next.Next)
		assert.Equal(t, node5, newHead.Next.Next.Next.Next)
	})

	t.Run("test case 2", func(t *testing.T) {
		node5 := &ListNode{Val: 5}
		node4 := &ListNode{Val: 4, Next: node5}
		node3 := &ListNode{Val: 3, Next: node4}
		node2 := &ListNode{Val: 2, Next: node3}
		node1 := &ListNode{Val: 1, Next: node2}
		k := 3
		newHead := reverseKGroup(node1, k)
		assert.Equal(t, node3, newHead)
		assert.Equal(t, node2, newHead.Next)
		assert.Equal(t, node1, newHead.Next.Next)
		assert.Equal(t, node4, newHead.Next.Next.Next)
		assert.Equal(t, node5, newHead.Next.Next.Next.Next)
	})
}
