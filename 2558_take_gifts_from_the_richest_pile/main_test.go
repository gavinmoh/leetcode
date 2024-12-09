package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPickGifts(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		gifts := []int{25, 64, 9, 4, 100}
		k := 4
		expected := int64(29)

		assert.Equal(t, expected, pickGifts(gifts, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		gifts := []int{1, 1, 1, 1}
		k := 4
		expected := int64(4)

		assert.Equal(t, expected, pickGifts(gifts, k))
	})
}
