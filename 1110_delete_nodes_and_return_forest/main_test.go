package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelNodes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
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
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		to_delete := []int{3, 5}
		expected := []*TreeNode{
			root,
			root.Right.Left,
			root.Right.Right,
		}

		assert.Equal(t, expected, delNodes(root, to_delete))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		}
		to_delete := []int{3}
		expected := []*TreeNode{
			root,
		}

		assert.Equal(t, expected, delNodes(root, to_delete))
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
					Val: 3,
				},
			},
		}
		to_delete := []int{2, 3}
		expected := []*TreeNode{
			root,
			root.Left.Left,
		}

		assert.Equal(t, expected, delNodes(root, to_delete))
	})
}
