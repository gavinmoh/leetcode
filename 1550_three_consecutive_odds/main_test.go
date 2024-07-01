package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeConsecutiveOdds(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{2, 6, 4, 1}

		assert.False(t, threeConsecutiveOdds(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}

		assert.True(t, threeConsecutiveOdds(arr))
	})
}
