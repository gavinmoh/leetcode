package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindThePrefixCommonArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		A := []int{1, 3, 2, 4}
		B := []int{3, 1, 2, 4}
		expected := []int{0, 2, 3, 4}

		assert.Equal(t, expected, findThePrefixCommonArray(A, B))
	})

	t.Run("test case 2", func(t *testing.T) {
		A := []int{2, 3, 1}
		B := []int{3, 1, 2}
		expected := []int{0, 1, 3}

		assert.Equal(t, expected, findThePrefixCommonArray(A, B))
	})
}
