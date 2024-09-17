package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUncommonFromSentences(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s1 := "this apple is sweet"
		s2 := "this apple is sour"
		expected := []string{"sweet", "sour"}

		assert.Equal(t, expected, uncommonFromSentences(s1, s2))
	})

	t.Run("test case 2", func(t *testing.T) {
		s1 := "apple apple"
		s2 := "banana"
		expected := []string{"banana"}

		assert.Equal(t, expected, uncommonFromSentences(s1, s2))
	})
}
