package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeAtIndex(t *testing.T) {
	case1 := "leet2code3"
	assert.Equal(t, "o", decodeAtIndex(case1, 10))

	case2 := "ha22"
	assert.Equal(t, "h", decodeAtIndex(case2, 5))

	case3 := "a2345678999999999999999"
	assert.Equal(t, "a", decodeAtIndex(case3, 1))

}
