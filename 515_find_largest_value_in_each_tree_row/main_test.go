package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestValues(t *testing.T) {
	t.Run("should return [1,3,9]", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 9,
				},
			},
		}

		expected := []int{1, 3, 9}
		assert.Equal(t, expected, largestValues(root))
	})

	t.Run("should return [1,3]", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}

		expected := []int{1, 3}
		assert.Equal(t, expected, largestValues(root))
	})

	t.Run("should return [-2147483648]", func(t *testing.T) {
		root := &TreeNode{
			Val: -2147483648,
		}

		expected := []int{-2147483648}
		assert.Equal(t, expected, largestValues(root))
	})
}
