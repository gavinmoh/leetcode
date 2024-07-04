package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeNodes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 0,
									},
								},
							},
						},
					},
				},
			},
		}
		expectedOutputs := []int{4, 11}
		head = mergeNodes(head)
		for _, expected := range expectedOutputs {
			assert.Equal(t, expected, head.Val)
			head = head.Next
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 0,
									},
								},
							},
						},
					},
				},
			},
		}
		expectedOutputs := []int{1, 3, 4}
		head = mergeNodes(head)
		for _, expected := range expectedOutputs {
			assert.Equal(t, expected, head.Val)
			head = head.Next
		}
	})
}
