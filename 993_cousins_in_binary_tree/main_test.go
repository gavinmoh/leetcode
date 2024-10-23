package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCousins(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		x := 4
		y := 3

		assert.False(t, isCousins(root, x, y))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		}
		x := 5
		y := 4

		assert.True(t, isCousins(root, x, y))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		x := 2
		y := 3

		assert.False(t, isCousins(root, x, y))
	})
}
