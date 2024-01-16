package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		node4 := &ListNode{Val: 4}
		node3 := &ListNode{Val: 3, Next: node4}
		node2 := &ListNode{Val: 2, Next: node3}
		node1 := &ListNode{Val: 1, Next: node2}
		newHead := swapPairs(node1)
		assert.Equal(t, node2, newHead)
		assert.Equal(t, node1, newHead.Next)
		assert.Equal(t, node4, newHead.Next.Next)
		assert.Equal(t, node3, newHead.Next.Next.Next)
	})

	t.Run("test case 2", func(t *testing.T) {
		var node1 *ListNode
		newHead := swapPairs(node1)
		assert.Equal(t, node1, newHead)
	})

	t.Run("test case 3", func(t *testing.T) {
		node1 := &ListNode{Val: 1}
		newHead := swapPairs(node1)
		assert.Equal(t, node1, newHead)
	})
}
