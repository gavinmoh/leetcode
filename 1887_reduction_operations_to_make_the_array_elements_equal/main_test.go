package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReductionOperations(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		nums := []int{5, 1, 3}
		expected := 3
		assert.Equal(t, expected, reductionOperations(nums))
	})

	t.Run("should return 0", func(t *testing.T) {
		nums := []int{1, 1, 1}
		expected := 0
		assert.Equal(t, expected, reductionOperations(nums))
	})

	t.Run("should return 4", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 3}
		expected := 4
		assert.Equal(t, expected, reductionOperations(nums))
	})

	t.Run("should return 1", func(t *testing.T) {
		nums := []int{2850, 2850, 2850, 2850, 2850, 2850, 2850, 2850, 2850, 20220, 2850}
		expected := 1
		assert.Equal(t, expected, reductionOperations(nums))
	})
}
