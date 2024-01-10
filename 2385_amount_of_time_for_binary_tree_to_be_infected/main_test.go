package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountOfTime(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		}
		start := 3
		expected := 4
		assert.Equal(t, expected, amountOfTime(root, start))
	})

	t.Run("should return 0", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		start := 1
		expected := 0
		assert.Equal(t, expected, amountOfTime(root, start))
	})
}
