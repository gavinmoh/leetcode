package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalanceBST(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
		}

		result := balanceBST(root)
		assert.Equal(t, 3, result.Val)
		assert.Equal(t, 2, result.Left.Val)
		assert.Equal(t, 1, result.Left.Left.Val)
		assert.Equal(t, 4, result.Right.Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}

		result := balanceBST(root)
		assert.Equal(t, 2, result.Val)
		assert.Equal(t, 1, result.Left.Val)
		assert.Equal(t, 3, result.Right.Val)
	})
}
