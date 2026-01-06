package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLevelSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: -8,
				},
			},
			Right: &TreeNode{
				Val: 0,
			},
		}
		expected := 2

		assert.Equal(t, expected, maxLevelSum(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 989,
			Right: &TreeNode{
				Val: 10250,
				Left: &TreeNode{
					Val: 98693,
				},
				Right: &TreeNode{
					Val: -89388,
					Right: &TreeNode{
						Val: -32127,
					},
				},
			},
		}
		expected := 2

		assert.Equal(t, expected, maxLevelSum(root))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: -100,
			Left: &TreeNode{
				Val: -200,
				Left: &TreeNode{
					Val: -20,
				},
				Right: &TreeNode{
					Val: -5,
				},
			},
			Right: &TreeNode{
				Val: -300,
				Left: &TreeNode{
					Val: -10,
				},
			},
		}
		expected := 3

		assert.Equal(t, expected, maxLevelSum(root))
	})
}
