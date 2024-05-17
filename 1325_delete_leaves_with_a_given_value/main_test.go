package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveLeafNodes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		}
		target := 2

		newRoot := removeLeafNodes(root, target)
		assert.Equal(t, 1, newRoot.Val)
		assert.Nil(t, newRoot.Left)
		assert.Equal(t, 3, newRoot.Right.Val)
		assert.Nil(t, newRoot.Right.Left)
		assert.Equal(t, 4, newRoot.Right.Right.Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		target := 3

		newRoot := removeLeafNodes(root, target)
		assert.Equal(t, 1, newRoot.Val)
		assert.Nil(t, newRoot.Right)
		assert.Equal(t, 3, newRoot.Left.Val)
		assert.Nil(t, newRoot.Left.Left)
		assert.Equal(t, 2, newRoot.Left.Right.Val)
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
		}
		target := 2

		newRoot := removeLeafNodes(root, target)
		assert.Equal(t, 1, newRoot.Val)
		assert.Nil(t, newRoot.Left)
		assert.Nil(t, newRoot.Right)
	})

	t.Run("test case 4", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 1,
			},
		}
		target := 1

		newRoot := removeLeafNodes(root, target)
		assert.Nil(t, newRoot)
	})

	t.Run("test case 5", func(t *testing.T) {
		root := &TreeNode{
			Val: 7,
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 6,
							},
							Right: &TreeNode{
								Val: 3,
							},
						},
					},
				},
			},
		}
		target := 3

		newRoot := removeLeafNodes(root, target)
		assert.Equal(t, 7, newRoot.Val)
		assert.Nil(t, newRoot.Left)
		assert.Equal(t, 6, newRoot.Right.Val)
		assert.Equal(t, 3, newRoot.Right.Left.Val)
		assert.Nil(t, newRoot.Right.Left.Left)
		assert.Equal(t, 2, newRoot.Right.Left.Right.Val)
		assert.Nil(t, newRoot.Right.Right.Left)
		assert.Equal(t, 3, newRoot.Right.Right.Val)
		assert.Nil(t, newRoot.Right.Right.Left)
		assert.Equal(t, 3, newRoot.Right.Right.Right.Val)
		assert.Nil(t, newRoot.Right.Right.Right.Right)
		assert.Equal(t, 3, newRoot.Right.Right.Right.Left.Val)
		assert.Nil(t, newRoot.Right.Right.Right.Left.Right)
		assert.Equal(t, 6, newRoot.Right.Right.Right.Left.Left.Val)
	})
}
