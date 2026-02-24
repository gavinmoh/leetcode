package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumRootToLeaf(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		}
		expected := 22

		assert.Equal(t, expected, sumRootToLeaf(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 0,
		}
		expected := 0

		assert.Equal(t, expected, sumRootToLeaf(root))
	})
}
