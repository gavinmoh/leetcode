package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree2str(t *testing.T) {
	t.Run("should return 1(2(4))(3)", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		expected := "1(2(4))(3)"
		assert.Equal(t, expected, tree2str(root))
	})

	t.Run("should return 1(2()(4))(3)", func(t *testing.T) {
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
		expected := "1(2()(4))(3)"
		assert.Equal(t, expected, tree2str(root))
	})

}
