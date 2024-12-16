package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseOddLevels(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 13,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 21,
				},
				Right: &TreeNode{
					Val: 34,
				},
			},
		}
		expected := []int{2, 5, 3, 8, 13, 21, 34}
		root = reverseOddLevels(root)

		assert.Equal(t, expected, treeToArr(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 11,
			},
		}
		expected := []int{7, 11, 13}
		root = reverseOddLevels(root)

		assert.Equal(t, expected, treeToArr(root))
	})

	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		}
		expected := []int{0, 2, 1, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1}
		root = reverseOddLevels(root)

		assert.Equal(t, expected, treeToArr(root))
	})
}

func treeToArr(root *TreeNode) []int {
	result := []int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		nextNodes := []*TreeNode{}
		for _, node := range nodes {
			if node == nil {
				continue
			}

			result = append(result, node.Val)
			nextNodes = append(nextNodes, node.Left, node.Right)
		}
		nodes = nextNodes
	}

	return result
}
