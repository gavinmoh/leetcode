package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifiedList(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3}
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
		expectedOutput := []int{4, 5}

		head = modifiedList(nums, head)
		for _, expected := range expectedOutput {
			assert.Equal(t, expected, head.Val)
			head = head.Next
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1}
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
							},
						},
					},
				},
			},
		}
		expectedOutput := []int{2, 2, 2}

		head = modifiedList(nums, head)
		for _, expected := range expectedOutput {
			assert.Equal(t, expected, head.Val)
			head = head.Next
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{5}
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
		expectedOutput := []int{1, 2, 3, 4}

		head = modifiedList(nums, head)
		for _, expected := range expectedOutput {
			assert.Equal(t, expected, head.Val)
			head = head.Next
		}
	})
}
