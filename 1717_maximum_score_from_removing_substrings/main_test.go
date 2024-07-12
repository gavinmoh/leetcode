package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumGain(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "cdbcbbaaabab"
		x := 4
		y := 5
		expected := 19

		assert.Equal(t, expected, maximumGain(s, x, y))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "aabbaaxybbaabb"
		x := 5
		y := 4
		expected := 20

		assert.Equal(t, expected, maximumGain(s, x, y))
	})

}
