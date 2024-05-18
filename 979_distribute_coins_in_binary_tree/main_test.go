package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCoins(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 0,
			},
		}
		expected := 2

		assert.Equal(t, expected, distributeCoins(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 0,
			},
		}
		expected := 3

		assert.Equal(t, expected, distributeCoins(root))
	})

}
