package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveZeroSumSublists(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: -3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		}
		expected := []int{3, 1}
		result := removeZeroSumSublists(head)

		for _, expectedValue := range expected {
			assert.Equal(t, expectedValue, result.Val)
			result = result.Next
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
						Val: -3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		}
		expected := []int{1, 2, 4}
		result := removeZeroSumSublists(head)

		for _, expectedValue := range expected {
			assert.Equal(t, expectedValue, result.Val)
			result = result.Next
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: -3,
						Next: &ListNode{
							Val: -2,
						},
					},
				},
			},
		}
		expected := []int{1}
		result := removeZeroSumSublists(head)

		for _, expectedValue := range expected {
			assert.Equal(t, expectedValue, result.Val)
			result = result.Next
		}
	})

	t.Run("test case 4", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: -1,
			},
		}
		assert.Nil(t, removeZeroSumSublists(head))
	})
}
