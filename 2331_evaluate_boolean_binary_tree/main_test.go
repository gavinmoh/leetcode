package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateTree(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		}

		assert.True(t, evaluateTree(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{Val: 0}

		assert.False(t, evaluateTree(root))
	})
}
