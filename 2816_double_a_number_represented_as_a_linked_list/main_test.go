package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleIt(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
				},
			},
		}
		output := []int{3, 7, 8}
		head = doubleIt(head)

		for _, expected := range output {
			assert.Equal(t, expected, head.Val)
			head = head.Next
		}
	})

	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		}
		output := []int{1, 9, 9, 8}
		head = doubleIt(head)

		for _, expected := range output {
			assert.Equal(t, expected, head.Val)
			head = head.Next
		}
	})
}
