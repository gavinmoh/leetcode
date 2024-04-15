package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumNumbers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		expected := 25

		assert.Equal(t, expected, sumNumbers(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 0,
			},
		}
		expected := 1026

		assert.Equal(t, expected, sumNumbers(root))
	})
}
