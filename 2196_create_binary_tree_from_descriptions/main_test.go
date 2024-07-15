package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBinaryTree(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		descriptions := [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}
		root := createBinaryTree(descriptions)

		assert.Equal(t, 50, root.Val)
		assert.Equal(t, 20, root.Left.Val)
		assert.Equal(t, 15, root.Left.Left.Val)
		assert.Equal(t, 17, root.Left.Right.Val)
		assert.Equal(t, 80, root.Right.Val)
		assert.Equal(t, 19, root.Right.Left.Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		descriptions := [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}}
		root := createBinaryTree(descriptions)

		assert.Equal(t, 1, root.Val)
		assert.Equal(t, 2, root.Left.Val)
		assert.Equal(t, 3, root.Left.Right.Val)
		assert.Equal(t, 4, root.Left.Right.Left.Val)
	})

}
