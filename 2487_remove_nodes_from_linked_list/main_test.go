package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNodes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 13,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 8,
						},
					},
				},
			},
		}
		output := []int{13, 8}
		current := removeNodes(head)

		for _, expected := range output {
			assert.Equal(t, expected, current.Val)
			current = current.Next
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		}
		output := []int{1, 1, 1, 1}
		current := removeNodes(head)

		for _, expected := range output {
			assert.Equal(t, expected, current.Val)
			current = current.Next
		}
	})
}
