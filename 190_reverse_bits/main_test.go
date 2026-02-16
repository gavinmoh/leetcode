package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 43261596
		expected := 964176192

		assert.Equal(t, expected, reverseBits(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 2147483644
		expected := 1073741822

		assert.Equal(t, expected, reverseBits(n))
	})
}
