package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlitEquiv(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root1 := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
			},
		}
		root2 := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 8,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		}

		assert.True(t, flipEquiv(root1, root2))
	})

	t.Run("test case 2", func(t *testing.T) {
		var root1 *TreeNode
		var root2 *TreeNode

		assert.True(t, flipEquiv(root1, root2))
	})

	t.Run("test case 3", func(t *testing.T) {
		var root1 *TreeNode
		var root2 *TreeNode = &TreeNode{Val: 1}

		assert.False(t, flipEquiv(root1, root2))
	})
}
