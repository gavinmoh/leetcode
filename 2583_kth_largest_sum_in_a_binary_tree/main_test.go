package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargestLevelSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		k := 2
		expected := int64(13)

		assert.Equal(t, expected, kthLargestLevelSum(root, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}
		k := 1
		expected := int64(3)

		assert.Equal(t, expected, kthLargestLevelSum(root, k))
	})

}
