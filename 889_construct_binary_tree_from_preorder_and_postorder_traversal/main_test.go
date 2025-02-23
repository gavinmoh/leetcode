package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructFromPrePost(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		preorder := []int{1, 2, 4, 5, 3, 6, 7}
		postorder := []int{4, 5, 2, 6, 7, 3, 1}
		root := constructFromPrePost(preorder, postorder)

		assert.Equal(t, 1, root.Val)
		assert.Equal(t, 2, root.Left.Val)
		assert.Equal(t, 4, root.Left.Left.Val)
		assert.Equal(t, 5, root.Left.Right.Val)
		assert.Equal(t, 3, root.Right.Val)
		assert.Equal(t, 6, root.Right.Left.Val)
		assert.Equal(t, 7, root.Right.Right.Val)
	})

	t.Run("test case 2", func(t *testing.T) {
		preorder := []int{1}
		postorder := []int{1}
		root := constructFromPrePost(preorder, postorder)

		assert.Equal(t, 1, root.Val)
	})

	t.Run("test case 3", func(t *testing.T) {
		preorder := []int{1, 2}
		postorder := []int{2, 1}
		root := constructFromPrePost(preorder, postorder)

		assert.Equal(t, 1, root.Val)
		assert.Equal(t, 2, root.Left.Val)
	})

	t.Run("test case 4", func(t *testing.T) {
		preorder := []int{2, 1, 3}
		postorder := []int{3, 1, 2}
		root := constructFromPrePost(preorder, postorder)

		assert.Equal(t, 2, root.Val)
		assert.Equal(t, 1, root.Left.Val)
		assert.Equal(t, 3, root.Left.Left.Val)
	})
}
