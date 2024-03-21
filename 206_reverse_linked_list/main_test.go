package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}
		expected := []int{5, 4, 3, 2, 1}
		result := reverseList(head)
		for _, num := range expected {
			assert.Equal(t, num, result.Val)
			result = result.Next
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		}
		expected := []int{2, 1}
		result := reverseList(head)
		for _, num := range expected {
			assert.Equal(t, num, result.Val)
			result = result.Next
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		var head *ListNode
		assert.Nil(t, reverseList(head))
	})
}
