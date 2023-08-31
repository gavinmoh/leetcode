package main

import "fmt"

func expandAroundCenter(s string, left int, right int) string {
	// left must be >= 0, since array starts at 0
	// right must be < len(s), since array ends at len(s) - 1
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	longest := ""

	for i, _ := range s {
		// we have to consider the substring palindrome can be
		// an odd or even length palindrome
		oddPalindrome := expandAroundCenter(s, i, i)
		evenPalindrome := expandAroundCenter(s, i, i+1)

		var currentLongest string
		if len(oddPalindrome) > len(evenPalindrome) {
			currentLongest = oddPalindrome
		} else {
			currentLongest = evenPalindrome
		}

		if len(currentLongest) > len(longest) {
			longest = currentLongest
		}
	}
	return longest
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
