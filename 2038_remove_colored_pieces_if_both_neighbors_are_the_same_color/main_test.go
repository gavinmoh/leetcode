package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWinnerOfGame(t *testing.T) {
	case1 := "AAABABB"
	assert.True(t, winnerOfGame(case1))

	case2 := "AA"
	assert.False(t, winnerOfGame(case2))

	case3 := "ABBBBBBBAAA"
	assert.False(t, winnerOfGame(case3))

}
