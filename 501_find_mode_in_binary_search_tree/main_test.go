package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMode(t *testing.T) {
	t.Run("should return [2]", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
			},
		}
		expected := []int{2}

		assert.Equal(t, expected, findMode(root))
	})

	t.Run("should return [0]", func(t *testing.T) {
		root := &TreeNode{
			Val: 0,
		}
		expected := []int{0}

		assert.Equal(t, expected, findMode(root))
	})

}
