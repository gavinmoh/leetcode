package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoverFromPreorder(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		traversal := "1-2--3--4-5--6--7"
		root := recoverFromPreorder(traversal)

		assert.Equal(t, root.Val, 1)
		assert.Equal(t, root.Left.Val, 2)
		assert.Equal(t, root.Left.Left.Val, 3)
		assert.Equal(t, root.Left.Right.Val, 4)
		assert.Equal(t, root.Right.Val, 5)
		assert.Equal(t, root.Right.Left.Val, 6)
		assert.Equal(t, root.Right.Right.Val, 7)
	})

	t.Run("test case 2", func(t *testing.T) {
		traversal := "1-2--3---4-5--6---7"
		root := recoverFromPreorder(traversal)

		assert.Equal(t, root.Val, 1)
		assert.Equal(t, root.Left.Val, 2)
		assert.Equal(t, root.Left.Left.Val, 3)
		assert.Equal(t, root.Left.Left.Left.Val, 4)
		assert.Equal(t, root.Right.Val, 5)
		assert.Equal(t, root.Right.Left.Val, 6)
		assert.Equal(t, root.Right.Left.Left.Val, 7)
	})

	t.Run("test case 3", func(t *testing.T) {
		traversal := "1-401--349---90--88"
		root := recoverFromPreorder(traversal)

		assert.Equal(t, root.Val, 1)
		assert.Equal(t, root.Left.Val, 401)
		assert.Equal(t, root.Left.Left.Val, 349)
		assert.Equal(t, root.Left.Left.Left.Val, 90)
		assert.Equal(t, root.Left.Right.Val, 88)
	})
}
