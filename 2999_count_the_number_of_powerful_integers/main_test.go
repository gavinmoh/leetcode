package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfPowerfulInt(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		start := int64(1)
		finish := int64(6000)
		limit := 4
		s := "124"
		expected := int64(5)

		assert.Equal(t, expected, numberOfPowerfulInt(start, finish, limit, s))
	})

	t.Run("test case 2", func(t *testing.T) {
		start := int64(15)
		finish := int64(215)
		limit := 6
		s := "10"
		expected := int64(2)

		assert.Equal(t, expected, numberOfPowerfulInt(start, finish, limit, s))
	})

	t.Run("test case 3", func(t *testing.T) {
		start := int64(1000)
		finish := int64(2000)
		limit := 4
		s := "3000"
		expected := int64(0)

		assert.Equal(t, expected, numberOfPowerfulInt(start, finish, limit, s))
	})
}
