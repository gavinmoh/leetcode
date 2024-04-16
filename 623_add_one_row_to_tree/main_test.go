package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOneRow(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 5,
				},
			},
		}
		val := 1
		depth := 2
		newRoot := addOneRow(root, val, depth)

		assert.Equal(t, 4, newRoot.Val)
		assert.Equal(t, 1, newRoot.Left.Val)
		assert.Equal(t, 2, newRoot.Left.Left.Val)
		assert.Equal(t, 1, newRoot.Right.Val)
		assert.Equal(t, 6, newRoot.Right.Right.Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		}
		val := 1
		depth := 3
		newRoot := addOneRow(root, val, depth)

		assert.Equal(t, 4, newRoot.Val)
		assert.Equal(t, 2, newRoot.Left.Val)
		assert.Equal(t, 1, newRoot.Left.Left.Val)
		assert.Equal(t, 3, newRoot.Left.Left.Left.Val)
		assert.Equal(t, 1, newRoot.Left.Right.Val)
		assert.Equal(t, 1, newRoot.Left.Right.Right.Val)
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 5,
				},
			},
		}
		val := 1
		depth := 1
		newRoot := addOneRow(root, val, depth)

		assert.Equal(t, 1, newRoot.Val)
		assert.Equal(t, 4, newRoot.Left.Val)
		assert.Equal(t, 2, newRoot.Left.Left.Val)
		assert.Equal(t, 6, newRoot.Left.Right.Val)
	})
}
