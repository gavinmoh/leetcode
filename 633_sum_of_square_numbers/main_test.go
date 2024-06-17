package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeSquareSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		c := 5
		assert.True(t, judgeSquareSum(c))
	})

	t.Run("test case 2", func(t *testing.T) {
		c := 3
		assert.False(t, judgeSquareSum(c))
	})

	t.Run("test case 3", func(t *testing.T) {
		c := 2
		assert.True(t, judgeSquareSum(c))
	})
}
