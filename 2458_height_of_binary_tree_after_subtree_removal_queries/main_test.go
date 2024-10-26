package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeQueries(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		}
		queries := []int{4}
		expected := []int{2}

		assert.Equal(t, expected, treeQueries(root, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		queries := []int{3, 2, 4, 8}
		expected := []int{3, 2, 3, 2}

		assert.Equal(t, expected, treeQueries(root, queries))
	})
}
