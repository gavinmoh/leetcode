package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralMatrix(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := arrayToLinkedList([]int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0})
		m := 3
		n := 5
		expected := [][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}}

		assert.Equal(t, expected, spiralMatrix(m, n, head))
	})

	t.Run("test case 2", func(t *testing.T) {
		head := arrayToLinkedList([]int{0, 1, 2})
		m := 1
		n := 4
		expected := [][]int{{0, 1, 2, -1}}

		assert.Equal(t, expected, spiralMatrix(m, n, head))
	})

	t.Run("test case 3", func(t *testing.T) {
		head := arrayToLinkedList([]int{8, 24, 5, 21, 10, 11, 11, 12, 6, 17})
		m := 10
		n := 1
		expected := [][]int{{8}, {24}, {5}, {21}, {10}, {11}, {11}, {12}, {6}, {17}}

		assert.Equal(t, expected, spiralMatrix(m, n, head))
	})
}

func arrayToLinkedList(array []int) *ListNode {
	head := &ListNode{Val: array[0]}
	current := head
	for i := 1; i < len(array); i++ {
		current.Next = &ListNode{Val: array[i]}
		current = current.Next
	}

	return head
}
