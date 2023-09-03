package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	case1 := 3
	assert.Equal(t, "III", IntToRoman(case1))

	case2 := 58
	assert.Equal(t, "LVIII", IntToRoman(case2))

	case3 := 1994
	assert.Equal(t, "MCMXCIV", IntToRoman(case3))
}
