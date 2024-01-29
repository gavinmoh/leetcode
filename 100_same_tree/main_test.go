package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
		q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

		assert.True(t, isSameTree(p, q))
	})

	t.Run("should return false", func(t *testing.T) {
		p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
		q := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

		assert.False(t, isSameTree(p, q))
	})
}
