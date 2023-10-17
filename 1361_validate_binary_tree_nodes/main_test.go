package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBinaryTreeNodes(t *testing.T) {
	case1LeftChild := []int{1, -1, 3, -1}
	case1RightChild := []int{2, -1, -1, -1}
	assert.True(t, validateBinaryTreeNodes(4, case1LeftChild, case1RightChild))

	case2LeftChild := []int{1, -1, 3, -1}
	case2RightChild := []int{2, 3, -1, -1}
	assert.False(t, validateBinaryTreeNodes(4, case2LeftChild, case2RightChild))

	case3LeftChild := []int{1, 0}
	case3RightChild := []int{-1, -1}
	assert.False(t, validateBinaryTreeNodes(2, case3LeftChild, case3RightChild))

	case4LeftChild := []int{3, -1, 1, -1}
	case4RightChild := []int{-1, -1, 0, -1}
	assert.True(t, validateBinaryTreeNodes(4, case4LeftChild, case4RightChild))

	case5LeftChild := []int{1, -1, -1, 4, -1, -1}
	case5RightChild := []int{2, -1, -1, 5, -1, -1}
	assert.False(t, validateBinaryTreeNodes(6, case5LeftChild, case5RightChild))
}
