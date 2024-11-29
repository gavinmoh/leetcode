package main

func canMakeSubsequence(str1 string, str2 string) bool {
	j := 0
	for i := 0; i < len(str1) && j < len(str2); i++ {
		if str1[i] == str2[j] || nextChar(str1[i]) == str2[j] {
			j++
		}
	}

	return j == len(str2)
}

func nextChar(char byte) byte {
	char++
	if char > 'z' {
		return 'a'
	}
	return char
}
