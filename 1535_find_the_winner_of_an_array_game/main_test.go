package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWinner(t *testing.T) {
	t.Run("should return 5", func(t *testing.T) {
		k := 2
		arr := []int{2, 1, 3, 5, 4, 6, 7}
		expected := 5
		assert.Equal(t, expected, getWinner(arr, k))
	})

	t.Run("should return 3", func(t *testing.T) {
		k := 10
		arr := []int{3, 2, 1}
		expected := 3
		assert.Equal(t, expected, getWinner(arr, k))
	})

}
