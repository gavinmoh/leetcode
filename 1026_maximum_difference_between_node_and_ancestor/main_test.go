package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAncestorDiff(t *testing.T) {
	t.Run("should return 7", func(t *testing.T) {
		root := &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			Right: &TreeNode{
				Val: 10,
				Right: &TreeNode{
					Val: 14,
					Left: &TreeNode{
						Val: 13,
					},
				},
			},
		}

		expected := 7
		assert.Equal(t, expected, maxAncestorDiff(root))
	})

	t.Run("should return 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		}

		expected := 3
		assert.Equal(t, expected, maxAncestorDiff(root))
	})

	t.Run("should return 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
				},
			},
		}

		expected := 2
		assert.Equal(t, expected, maxAncestorDiff(root))
	})
}
