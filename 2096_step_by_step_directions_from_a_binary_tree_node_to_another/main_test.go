package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirections(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		}
		startValue := 3
		destValue := 6
		expected := "UURL"

		assert.Equal(t, expected, getDirections(root, startValue, destValue))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		}
		startValue := 2
		destValue := 1
		expected := "L"

		assert.Equal(t, expected, getDirections(root, startValue, destValue))
	})
}
