package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
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
		expected := head.Next.Next

		assert.Equal(t, expected, middleNode(head))
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
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		}
		expected := head.Next.Next.Next

		assert.Equal(t, expected, middleNode(head))
	})
}
