package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeInBetween(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		list1 := &ListNode{
			Val: 10,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 13,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		}
		list2 := &ListNode{
			Val: 1000000,
			Next: &ListNode{
				Val: 1000001,
				Next: &ListNode{
					Val: 1000002,
				},
			},
		}
		a := 3
		b := 4
		expected := []int{10, 1, 13, 1000000, 1000001, 1000002, 5}
		result := mergeInBetween(list1, a, b, list2)
		for _, num := range expected {
			assert.Equal(t, num, result.Val)
			result = result.Next
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		list1 := &ListNode{
			Val: 0,
			Next: &ListNode{
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
			},
		}
		list2 := &ListNode{
			Val: 1000000,
			Next: &ListNode{
				Val: 1000001,
				Next: &ListNode{
					Val: 1000002,
					Next: &ListNode{
						Val: 1000003,
						Next: &ListNode{
							Val: 1000004,
						},
					},
				},
			},
		}
		a := 2
		b := 5
		expected := []int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}
		result := mergeInBetween(list1, a, b, list2)
		for _, num := range expected {
			assert.Equal(t, num, result.Val)
			result = result.Next
		}
	})
}
