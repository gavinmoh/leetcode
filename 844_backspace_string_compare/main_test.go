package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackspaceCompare(t *testing.T) {
	case1S := "ab#c"
	case1T := "ad#c"
	assert.True(t, backspaceCompare(case1S, case1T))

	case2S := "ab##"
	case2T := "c#d#"
	assert.True(t, backspaceCompare(case2S, case2T))

	case3S := "a##c"
	case3T := "b"
	assert.False(t, backspaceCompare(case3S, case3T))
}
