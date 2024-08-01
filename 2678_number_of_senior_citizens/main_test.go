package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSeniors(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		details := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
		expected := 2

		assert.Equal(t, expected, countSeniors(details))
	})

	t.Run("test case 2", func(t *testing.T) {
		details := []string{"1313579440F2036", "2921522980M5644"}
		expected := 0

		assert.Equal(t, expected, countSeniors(details))
	})
}
