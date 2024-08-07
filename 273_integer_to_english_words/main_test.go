package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToWords(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num := 100
		expected := "One Hundred"

		assert.Equal(t, expected, numberToWords(num))
	})

	t.Run("test case 2", func(t *testing.T) {
		num := 12_345
		expected := "Twelve Thousand Three Hundred Forty Five"

		assert.Equal(t, expected, numberToWords(num))
	})

	t.Run("test case 3", func(t *testing.T) {
		num := 1_234_567
		expected := "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

		assert.Equal(t, expected, numberToWords(num))
	})

	t.Run("test case 4", func(t *testing.T) {
		num := 1_234_567_891
		expected := "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"

		assert.Equal(t, expected, numberToWords(num))
	})

	t.Run("test case 5", func(t *testing.T) {
		num := 30
		expected := "Thirty"

		assert.Equal(t, expected, numberToWords(num))
	})

	t.Run("test case 6", func(t *testing.T) {
		num := 1_000
		expected := "One Thousand"

		assert.Equal(t, expected, numberToWords(num))
	})
}
