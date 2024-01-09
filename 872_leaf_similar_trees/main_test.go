package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeafSimilar(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		root1 := &TreeNode{
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
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		}

		root2 := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		}

		assert.True(t, leafSimilar(root1, root2))
	})

	t.Run("should return false", func(t *testing.T) {
		root1 := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}

		root2 := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 2,
			},
		}

		assert.False(t, leafSimilar(root1, root2))
	})
}
