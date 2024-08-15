package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonandeChange(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		bills := []int{5, 5, 5, 10, 20}

		assert.True(t, lemonadeChange(bills))
	})

	t.Run("test case 2", func(t *testing.T) {
		bills := []int{5, 5, 10, 10, 20}

		assert.False(t, lemonadeChange(bills))
	})
}
