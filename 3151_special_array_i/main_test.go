package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecialArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1}

		assert.True(t, isArraySpecial(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 1, 4}

		assert.True(t, isArraySpecial(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{4, 3, 1, 6}

		assert.False(t, isArraySpecial(nums))
	})
}
