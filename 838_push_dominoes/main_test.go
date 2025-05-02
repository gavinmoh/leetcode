package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushDominoes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		dominoes := "RR.L"
		expected := "RR.L"

		assert.Equal(t, expected, pushDominoes(dominoes))
	})

	t.Run("test case 2", func(t *testing.T) {
		dominoes := ".L.R...LR..L.."
		expected := "LL.RR.LLRRLL.."

		assert.Equal(t, expected, pushDominoes(dominoes))
	})
}
