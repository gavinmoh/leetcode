package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceNonCoprimes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{6, 4, 3, 2, 7, 6, 2}
		expected := []int{12, 7, 6}

		assert.Equal(t, expected, replaceNonCoprimes(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 2, 1, 1, 3, 3, 3}
		expected := []int{2, 1, 1, 3}

		assert.Equal(t, expected, replaceNonCoprimes(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{287, 41, 49, 287, 899, 23, 23, 20677, 5, 825}
		expected := []int{2009, 20677, 825}

		assert.Equal(t, expected, replaceNonCoprimes(nums))
	})
}
