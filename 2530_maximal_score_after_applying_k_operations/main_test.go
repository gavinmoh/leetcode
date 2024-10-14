package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxKelements(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{10, 10, 10, 10, 10}
		k := 5
		expected := int64(50)

		assert.Equal(t, expected, maxKelements(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 10, 3, 3, 3}
		k := 3
		expected := int64(17)

		assert.Equal(t, expected, maxKelements(nums, k))
	})
}
