package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEvenOddTree(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		}
		assert.True(t, isEvenOddTree(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
			},
		}
		assert.False(t, isEvenOddTree(root))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 7,
				},
			},
		}
		assert.False(t, isEvenOddTree(root))
	})

	t.Run("test case 4", func(t *testing.T) {
		root := &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 30,
						Left: &TreeNode{
							Val: 17,
						},
					},
					Right: &TreeNode{
						Val: 20,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 18,
					},
					Right: &TreeNode{
						Val: 16,
					},
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 10,
					},
				},
				Right: &TreeNode{
					Val: 11,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		}
		assert.True(t, isEvenOddTree(root))
	})
}
