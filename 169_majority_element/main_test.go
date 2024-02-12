package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		nums := []int{3, 2, 3}
		expected := 3

		assert.Equal(t, expected, majorityElement(nums))
	})

	t.Run("should return 2", func(t *testing.T) {
		nums := []int{2, 2, 1, 1, 1, 2, 2}
		expected := 2

		assert.Equal(t, expected, majorityElement(nums))
	})

}
