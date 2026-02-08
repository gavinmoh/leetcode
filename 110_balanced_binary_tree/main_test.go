package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}

		assert.True(t, isBalanced(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 2},
		}

		assert.False(t, isBalanced(root))
	})

	t.Run("test case 3", func(t *testing.T) {
		var root *TreeNode

		assert.True(t, isBalanced(root))
	})

	t.Run("test case 4", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}

		assert.True(t, isBalanced(root))
	})

	t.Run("test case 5", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		}

		assert.False(t, isBalanced(root))
	})
}
