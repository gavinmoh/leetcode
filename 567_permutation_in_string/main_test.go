package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s1 := "ab"
		s2 := "eidbaooo"

		assert.True(t, checkInclusion(s1, s2))
	})

	t.Run("test case 2", func(t *testing.T) {
		s1 := "ab"
		s2 := "eidboaoo"

		assert.False(t, checkInclusion(s1, s2))
	})

	t.Run("test case 3", func(t *testing.T) {
		s1 := "adc"
		s2 := "dcda"

		assert.True(t, checkInclusion(s1, s2))
	})
}
