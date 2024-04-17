package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestFromLeaf(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		}
		expected := "dba"

		assert.Equal(t, expected, smallestFromLeaf(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 25,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		}
		expected := "adz"

		assert.Equal(t, expected, smallestFromLeaf(root))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
			},
		}
		expected := "abc"

		assert.Equal(t, expected, smallestFromLeaf(root))
	})

	t.Run("test case 4", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		expected := "hud"

		assert.Equal(t, expected, smallestFromLeaf(root))
	})
}
