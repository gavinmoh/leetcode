package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRecord(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "PPALLP"

		assert.True(t, checkRecord(s))
	})

	t.Run("test case 1", func(t *testing.T) {
		s := "PPALLL"

		assert.False(t, checkRecord(s))
	})
}
