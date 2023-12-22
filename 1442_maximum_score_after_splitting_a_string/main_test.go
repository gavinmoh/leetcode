package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScore(t *testing.T) {
	t.Run("should return 5", func(t *testing.T) {
		s := "011101"
		expected := 5
		assert.Equal(t, expected, maxScore(s))
	})

	t.Run("should return 5", func(t *testing.T) {
		s := "00111"
		expected := 5
		assert.Equal(t, expected, maxScore(s))
	})

	t.Run("should return 3", func(t *testing.T) {
		s := "1111"
		expected := 3
		assert.Equal(t, expected, maxScore(s))
	})

}
