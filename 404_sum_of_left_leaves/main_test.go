package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfLeftLeaves(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		expected := 24

		assert.Equal(t, expected, sumOfLeftLeaves(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
		}
		expected := 0

		assert.Equal(t, expected, sumOfLeftLeaves(root))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 20,
			},
		}
		expected := 4

		assert.Equal(t, expected, sumOfLeftLeaves(root))
	})

	t.Run("test case 4", func(t *testing.T) {
		root := &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: -1,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		}
		expected := 5

		assert.Equal(t, expected, sumOfLeftLeaves(root))
	})
}
