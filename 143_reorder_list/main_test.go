package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		}
		expected := []int{1, 4, 2, 3}
		reorderList(head)

		for _, num := range expected {
			assert.Equal(t, num, head.Val)
			head = head.Next
		}
	})

	t.Run("test case 2", func(t *testing.T) {
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
		expected := []int{1, 5, 2, 4, 3}
		reorderList(head)

		for _, num := range expected {
			assert.Equal(t, num, head.Val)
			head = head.Next
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		}
		expected := []int{1, 3, 2}
		reorderList(head)

		for _, num := range expected {
			assert.Equal(t, num, head.Val)
			head = head.Next
		}
	})
}
