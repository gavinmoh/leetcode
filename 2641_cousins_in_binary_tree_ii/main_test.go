package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceValueInTree(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 10,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		root = replaceValueInTree(root)

		assert.Equal(t, 0, root.Val)
		assert.Equal(t, 0, root.Left.Val)
		assert.Equal(t, 0, root.Right.Val)
		assert.Equal(t, 7, root.Left.Left.Val)
		assert.Equal(t, 7, root.Left.Right.Val)
		assert.Equal(t, 11, root.Right.Right.Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		}
		root = replaceValueInTree(root)

		assert.Equal(t, 0, root.Val)
		assert.Equal(t, 0, root.Left.Val)
		assert.Equal(t, 0, root.Right.Val)
	})
}
