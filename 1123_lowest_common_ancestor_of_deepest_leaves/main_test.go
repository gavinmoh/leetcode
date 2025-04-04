package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLcaDeepestLeaves(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		}
		expected := 2

		assert.Equal(t, expected, lcaDeepestLeaves(root).Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		expected := 1

		assert.Equal(t, expected, lcaDeepestLeaves(root).Val)
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		expected := 2

		assert.Equal(t, expected, lcaDeepestLeaves(root).Val)
	})
}
