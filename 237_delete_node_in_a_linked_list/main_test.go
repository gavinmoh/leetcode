package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		}
		output := []int{4, 1, 9}

		deleteNode(head.Next)

		current := head
		for _, expected := range output {
			assert.Equal(t, expected, current.Val)
			current = current.Next
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		}
		output := []int{4, 5, 9}

		deleteNode(head.Next.Next)

		current := head
		for _, expected := range output {
			assert.Equal(t, expected, current.Val)
			current = current.Next
		}
	})
}
