package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 9,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 10,
					},
				},
			},
		}
		expected := 3

		assert.Equal(t, expected, minimumOperations(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		}
		expected := 3

		assert.Equal(t, expected, minimumOperations(root))
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
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
			},
		}
		expected := 0

		assert.Equal(t, expected, minimumOperations(root))
	})

	t.Run("test case 4", func(t *testing.T) {
		root := &TreeNode{
			Val: 49,
			Left: &TreeNode{
				Val: 45,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 27,
					},
				},
				Right: &TreeNode{
					Val: 46,
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 25,
					},
				},
				Right: &TreeNode{
					Val: 39,
				},
			},
		}
		expected := 5

		assert.Equal(t, expected, minimumOperations(root))
	})
}
