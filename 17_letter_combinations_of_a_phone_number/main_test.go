package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	case1 := "23"
	case1Expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	assert.Equal(t, case1Expected, LetterCombinations(case1))

	case2 := ""
	case2Expected := []string{}
	assert.Equal(t, case2Expected, LetterCombinations(case2))

	case3 := "2"
	case3Expected := []string{"a", "b", "c"}
	assert.Equal(t, case3Expected, LetterCombinations(case3))
}
