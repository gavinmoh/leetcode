package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindErrorNums(t *testing.T) {
	t.Run("should return [2,3]", func(t *testing.T) {
		nums := []int{1, 2, 2, 4}
		expected := []int{2, 3}
		assert.Equal(t, expected, findErrorNums(nums))
	})

	t.Run("should return [1,2]", func(t *testing.T) {
		nums := []int{1, 1}
		expected := []int{1, 2}
		assert.Equal(t, expected, findErrorNums(nums))
	})

	t.Run("should return [2,1]", func(t *testing.T) {
		nums := []int{2, 2}
		expected := []int{2, 1}
		assert.Equal(t, expected, findErrorNums(nums))
	})
}
