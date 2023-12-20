package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyChoco(t *testing.T) {
	t.Run("should return 0", func(t *testing.T) {
		prices := []int{1, 2, 2}
		money := 3
		expected := 0
		assert.Equal(t, expected, buyChoco(prices, money))
	})

	t.Run("should return 3", func(t *testing.T) {
		prices := []int{3, 2, 3}
		money := 3
		expected := 3
		assert.Equal(t, expected, buyChoco(prices, money))
	})
}
