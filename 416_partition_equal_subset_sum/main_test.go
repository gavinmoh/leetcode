package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 5, 11, 5}

		assert.True(t, canPartition(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 5}

		assert.False(t, canPartition(nums))
	})
}
