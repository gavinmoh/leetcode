package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 4
	expected := "PINALSIGYAHRPI"
	assert.Equal(t, expected, convert(s, numRows))
}
