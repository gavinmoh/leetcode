package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinExtraChar(t *testing.T) {
	case1s := "leetscode"
	case1dictionary := []string{"leet", "code", "leetcode"}
	assert.Equal(t, 1, minExtraChar(case1s, case1dictionary))

	case2s := "sayhelloworld"
	case2dictionary := []string{"hello", "world"}
	assert.Equal(t, 3, minExtraChar(case2s, case2dictionary))

	case3s := "metzeaencgpgvsckjrqafkxgyzbe"
	case3dictionary := []string{"zdzz", "lgrhy", "r", "ohk", "zkowk", "g", "zqpn", "anoni", "ka", "qafkx", "t", "jr", "xdye", "mppc", "bqqb", "encgp", "yf", "vl", "ctsxk", "gn", "cujh", "ce", "rwrpq", "tze", "zxhg", "yzbe", "c", "o", "hnk", "gv", "uzbc", "xn", "kk", "ujjd", "vv", "mxhmv", "ugn", "at", "kumr", "ensv", "x", "uy", "gb", "ae", "jljuo", "xqkgj"}
	assert.Equal(t, 5, minExtraChar(case3s, case3dictionary))
}
