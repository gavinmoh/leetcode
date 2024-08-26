package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostOrder(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &Node{
			Val: 1,
			Children: []*Node{
				&Node{
					Val: 3,
					Children: []*Node{
						&Node{Val: 5},
						&Node{Val: 6},
					},
				},
				&Node{Val: 2},
				&Node{Val: 4},
			},
		}
		expected := []int{5, 6, 3, 2, 4, 1}

		assert.Equal(t, expected, postorder(root))
	})

	t.Run("test case 2", func(t *testing.T) {
		root := &Node{
			Val: 1,
			Children: []*Node{
				&Node{Val: 2},
				&Node{
					Val: 3,
					Children: []*Node{
						&Node{Val: 6},
						&Node{
							Val: 7,
							Children: []*Node{
								&Node{
									Val: 11,
									Children: []*Node{
										&Node{Val: 14},
									},
								},
							},
						},
					},
				},
				&Node{
					Val: 4,
					Children: []*Node{
						&Node{
							Val: 8,
							Children: []*Node{
								&Node{Val: 12},
							},
						},
					},
				},
				&Node{
					Val: 5,
					Children: []*Node{
						&Node{
							Val: 9,
							Children: []*Node{
								&Node{Val: 13},
							},
						},
						&Node{Val: 10},
					},
				},
			},
		}
		expected := []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1}

		assert.Equal(t, expected, postorder(root))
	})
}
