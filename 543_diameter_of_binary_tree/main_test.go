package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
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
			},
		}
		expected := 3

		assert.Equal(t, expected, diameterOfBinaryTree(root))
	})

	t.Run("should return 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}
		expected := 1

		assert.Equal(t, expected, diameterOfBinaryTree(root))
	})

	t.Run("should return 8", func(t *testing.T) {
		root := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: -7,
			},
			Right: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val: -9,
					Left: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 0,
								Right: &TreeNode{
									Val: -1,
								},
							},
							Right: &TreeNode{
								Val: 6,
								Left: &TreeNode{
									Val: -4,
								},
							},
						},
					},
					Right: &TreeNode{
						Val: -7,
						Left: &TreeNode{
							Val: -6,
							Left: &TreeNode{
								Val: 5,
							},
						},
						Right: &TreeNode{
							Val: -6,
							Left: &TreeNode{
								Val: 9,
								Left: &TreeNode{
									Val: -2,
								},
							},
						},
					},
				},
				Right: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -4,
					},
				},
			},
		}
		expected := 8

		assert.Equal(t, expected, diameterOfBinaryTree(root))
	})
}
