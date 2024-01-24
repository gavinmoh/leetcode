package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPseudoPalindromicPaths(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
				},
			},
		}
		expected := 2
		assert.Equal(t, expected, pseudoPalindromicPaths(root))
	})

	t.Run("should return 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
			},
		}
		expected := 1
		assert.Equal(t, expected, pseudoPalindromicPaths(root))
	})

	t.Run("should return 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 9,
		}
		expected := 1
		assert.Equal(t, expected, pseudoPalindromicPaths(root))
	})
}
