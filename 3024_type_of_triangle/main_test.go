package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangleType(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 3, 3}
		expected := "equilateral"

		assert.Equal(t, expected, triangleType(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 4, 5}
		expected := "scalene"

		assert.Equal(t, expected, triangleType(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{9, 4, 9}
		expected := "isosceles"

		assert.Equal(t, expected, triangleType(nums))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{8, 4, 4}
		expected := "none"

		assert.Equal(t, expected, triangleType(nums))
	})
}
