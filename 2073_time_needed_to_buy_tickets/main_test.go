package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeRequiredToBuy(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		tickets := []int{2, 3, 2}
		k := 2
		expected := 6

		assert.Equal(t, expected, timeRequiredToBuy(tickets, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		tickets := []int{5, 1, 1, 1}
		k := 0
		expected := 8

		assert.Equal(t, expected, timeRequiredToBuy(tickets, k))
	})
}
