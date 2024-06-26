package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBstToGst(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		}

		bstToGst(root)

		assert.Equal(t, 30, root.Val)
		assert.Equal(t, 36, root.Left.Val)
		assert.Equal(t, 36, root.Left.Left.Val)
		assert.Equal(t, 35, root.Left.Right.Val)
		assert.Equal(t, 33, root.Left.Right.Right.Val)
		assert.Equal(t, 21, root.Right.Val)
		assert.Equal(t, 26, root.Right.Left.Val)
		assert.Equal(t, 15, root.Right.Right.Val)
		assert.Equal(t, 8, root.Right.Right.Right.Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val: 1,
			},
		}

		bstToGst(root)

		assert.Equal(t, 1, root.Val)
		assert.Equal(t, 1, root.Right.Val)
	})
}
