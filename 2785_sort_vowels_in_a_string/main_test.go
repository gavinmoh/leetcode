package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortVowels(t *testing.T) {
	t.Run("should return lEOtcede", func(t *testing.T) {
		s := "lEetcOde"
		expected := "lEOtcede"
		assert.Equal(t, expected, sortVowels(s))
	})

	t.Run("should return lYmpH", func(t *testing.T) {
		s := "lYmpH"
		expected := "lYmpH"
		assert.Equal(t, expected, sortVowels(s))
	})

}
