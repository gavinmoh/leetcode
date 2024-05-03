package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	for i := 0; i < max(len(v1), len(v2)); i++ {
		v1Int, v2Int := 0, 0
		if i < len(v1) {
			v1Int, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			v2Int, _ = strconv.Atoi(v2[i])
		}

		if v1Int > v2Int {
			return 1
		} else if v2Int > v1Int {
			return -1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
