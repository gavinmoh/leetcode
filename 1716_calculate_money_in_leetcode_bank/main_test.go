package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalMoney(t *testing.T) {
	t.Run("should return 10", func(t *testing.T) {
		n := 4
		expected := 10
		assert.Equal(t, expected, totalMoney(n))
	})

	t.Run("should return 37", func(t *testing.T) {
		n := 10
		expected := 37
		assert.Equal(t, expected, totalMoney(n))
	})

}
