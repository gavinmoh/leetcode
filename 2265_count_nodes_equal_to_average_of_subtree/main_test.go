package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageOfSubtree(t *testing.T) {
	t.Run("should return 5", func(t *testing.T) {
		root := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 6,
				},
			},
		}
		expected := 5
		assert.Equal(t, expected, averageOfSubtree(root))
	})

	t.Run("should return 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
		}
		expected := 1
		assert.Equal(t, expected, averageOfSubtree(root))
	})
}
