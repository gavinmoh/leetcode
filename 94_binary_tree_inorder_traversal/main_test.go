package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	t.Run("should return [1,3,2]", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}
		expected := []int{1, 3, 2}
		assert.Equal(t, expected, inorderTraversal(root))
	})
	t.Run("should return []", func(t *testing.T) {
		var root *TreeNode
		expected := []int{}
		assert.Equal(t, expected, inorderTraversal(root))
	})

	t.Run("should return [1]", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		expected := []int{1}
		assert.Equal(t, expected, inorderTraversal(root))
	})
}
