package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		distance := 3
		expected := 1

		assert.Equal(t, expected, countPairs(root, distance))
	})

	t.Run("test case 2", func(t *testing.T) {
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
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		distance := 3
		expected := 2

		assert.Equal(t, expected, countPairs(root, distance))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		}
		distance := 3
		expected := 1

		assert.Equal(t, expected, countPairs(root, distance))
	})
}
