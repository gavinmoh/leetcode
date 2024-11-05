package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeSubOperation(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{4, 9, 6, 10}

		assert.True(t, primeSubOperation(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{6, 8, 11, 12}

		assert.True(t, primeSubOperation(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{5, 8, 3}

		assert.False(t, primeSubOperation(nums))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{998, 2}

		assert.True(t, primeSubOperation(nums))
	})

	t.Run("test case 5", func(t *testing.T) {
		nums := []int{2, 2}

		assert.False(t, primeSubOperation(nums))
	})

	t.Run("test case 6", func(t *testing.T) {
		nums := []int{19, 10}

		assert.True(t, primeSubOperation(nums))
	})

	t.Run("test case 7", func(t *testing.T) {
		nums := []int{15, 20, 17, 7, 16}

		assert.True(t, primeSubOperation(nums))
	})

	t.Run("test case 8", func(t *testing.T) {
		nums := []int{5, 16, 14, 9}

		assert.True(t, primeSubOperation(nums))
	})
}
