package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}
		expected := []int{3, 2, 1}

		assert.Equal(t, expected, postorderTraversal(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		var root *TreeNode
		expected := []int{}

		assert.Equal(t, expected, postorderTraversal(root))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		expected := []int{1}

		assert.Equal(t, expected, postorderTraversal(root))
	})
}
