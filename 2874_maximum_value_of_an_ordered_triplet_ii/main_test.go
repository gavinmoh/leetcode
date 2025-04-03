package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumTripletValue(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{12, 6, 1, 2, 7}
		expected := int64(77)

		assert.Equal(t, expected, maximumTripletValue(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 10, 3, 4, 19}
		expected := int64(133)

		assert.Equal(t, expected, maximumTripletValue(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := int64(0)

		assert.Equal(t, expected, maximumTripletValue(nums))
	})
}
