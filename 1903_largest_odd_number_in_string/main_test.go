package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestOddNumber(t *testing.T) {
	t.Run("should return 5", func(t *testing.T) {
		num := "52"
		expected := "5"
		assert.Equal(t, expected, largestOddNumber(num))
	})

	t.Run("should return empty string", func(t *testing.T) {
		num := "4206"
		expected := ""
		assert.Equal(t, expected, largestOddNumber(num))
	})

	t.Run("should return empty 35427", func(t *testing.T) {
		num := "35427"
		expected := "35427"
		assert.Equal(t, expected, largestOddNumber(num))
	})

}
