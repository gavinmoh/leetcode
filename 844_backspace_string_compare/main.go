package main

func backspaceCompare(s string, t string) bool {

	sStr := []rune{}
	tStr := []rune{}

	for _, char := range s {
		if char == '#' {
			if len(sStr) > 0 {
				sStr = sStr[:len(sStr)-1] // remove the last character
			}
		} else {
			sStr = append(sStr, char) // add the character
		}
	}

	for _, char := range t {
		if char == '#' {
			if len(tStr) > 0 {
				tStr = tStr[:len(tStr)-1] // remove the last character
			}
		} else {
			tStr = append(tStr, char) // add the character
		}
	}

	return string(sStr) == string(tStr)
}
