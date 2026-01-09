package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
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
		expected := 3

		assert.Equal(t, expected, maxDepth(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		}
		expected := 2

		assert.Equal(t, expected, maxDepth(root))
	})

	t.Run("test case 3", func(t *testing.T) {
		var root *TreeNode
		expected := 0

		assert.Equal(t, expected, maxDepth(root))
	})
}
