package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindElements(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: -1,
			Right: &TreeNode{
				Val: -1,
			},
		}
		obj := Constructor(root)

		assert.False(t, obj.Find(1))
		assert.True(t, obj.Find(2))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: -1,
			Left: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: -1,
				},
				Right: &TreeNode{
					Val: -1,
				},
			},
			Right: &TreeNode{
				Val: -1,
			},
		}
		obj := Constructor(root)

		assert.True(t, obj.Find(1))
		assert.True(t, obj.Find(3))
		assert.False(t, obj.Find(5))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: -1,
			Right: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: -1,
					Left: &TreeNode{
						Val: -1,
					},
				},
			},
		}
		obj := Constructor(root)

		assert.True(t, obj.Find(2))
		assert.False(t, obj.Find(3))
		assert.False(t, obj.Find(4))
		assert.True(t, obj.Find(5))
	})
}
