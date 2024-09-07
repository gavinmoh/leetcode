package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubPath(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 8,
				},
			},
		}
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
		}

		assert.True(t, isSubPath(head, root))
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		}
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
		}

		assert.True(t, isSubPath(head, root))
	})

	t.Run("test case 3", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 8,
						},
					},
				},
			},
		}
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
		}

		assert.False(t, isSubPath(head, root))
	})
}
