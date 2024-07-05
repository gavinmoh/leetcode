package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodesBetweenCriticalPoints(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 1,
			},
		}
		expected := []int{-1, -1}

		assert.Equal(t, expected, nodesBetweenCriticalPoints(head))
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val: 2,
								},
							},
						},
					},
				},
			},
		}
		expected := []int{1, 3}

		assert.Equal(t, expected, nodesBetweenCriticalPoints(head))
	})

	t.Run("test case 3", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 2,
										Next: &ListNode{
											Val: 7,
										},
									},
								},
							},
						},
					},
				},
			},
		}
		expected := []int{3, 3}

		assert.Equal(t, expected, nodesBetweenCriticalPoints(head))
	})

	t.Run("test case 4", func(t *testing.T) {
		head := &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		}
		expected := []int{-1, -1}

		assert.Equal(t, expected, nodesBetweenCriticalPoints(head))
	})
}
