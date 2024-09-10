package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertGreatestCommonDivisors(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := arrayToList([]int{18, 6, 10, 3})
		expected := []int{18, 6, 6, 2, 10, 1, 3}

		assert.Equal(t, expected, listToArray(insertGreatestCommonDivisors(head)))
	})

	t.Run("test case 2", func(t *testing.T) {
		head := arrayToList([]int{7})
		expected := []int{7}

		assert.Equal(t, expected, listToArray(insertGreatestCommonDivisors(head)))
	})
}

func arrayToList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

func listToArray(node *ListNode) []int {
	result := []int{}
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	return result
}
