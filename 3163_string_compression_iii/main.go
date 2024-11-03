package main

import (
	"fmt"
	"strings"
)

func compressedString(word string) string {
	var comp strings.Builder
	count := 1

	for i := 1; i < len(word); i++ {
		if word[i] != word[i-1] || count == 9 {
			comp.WriteString(fmt.Sprintf("%d%c", count, word[i-1]))
			count = 1
		} else {
			count += 1
		}
	}
	comp.WriteString(fmt.Sprintf("%d%c", count, word[len(word)-1]))

	return comp.String()
}
