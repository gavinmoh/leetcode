package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatLimitedString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "cczazcc"
		repeatLimit := 3
		expected := "zzcccac"

		assert.Equal(t, expected, repeatLimitedString(s, repeatLimit))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "aababab"
		repeatLimit := 2
		expected := "bbabaa"

		assert.Equal(t, expected, repeatLimitedString(s, repeatLimit))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "robnsdvpuxbapuqgopqvxdrchivlifeepy"
		repeatLimit := 2
		expected := "yxxvvuvusrrqqppopponliihgfeeddcbba"

		assert.Equal(t, expected, repeatLimitedString(s, repeatLimit))
	})
}
