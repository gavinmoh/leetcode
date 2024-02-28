package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBottomLeftValue(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		expected := 1

		assert.Equal(t, expected, findBottomLeftValue(root))
	})

	t.Run("should return 7", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		}
		expected := 7

		assert.Equal(t, expected, findBottomLeftValue(root))
	})
}
