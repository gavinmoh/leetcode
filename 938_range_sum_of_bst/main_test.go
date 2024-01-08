package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSumBT(t *testing.T) {
	t.Run("should return 32", func(t *testing.T) {
		root := &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Right: &TreeNode{
					Val: 18,
				},
			},
		}
		low := 7
		high := 15
		expected := 32
		assert.Equal(t, expected, rangeSumBST(root, low, high))
	})

	t.Run("should return 23", func(t *testing.T) {
		root := &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
				},
				Right: &TreeNode{
					Val: 18,
				},
			},
		}
		low := 6
		high := 10
		expected := 23
		assert.Equal(t, expected, rangeSumBST(root, low, high))
	})
}
