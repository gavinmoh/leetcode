package main

func isCircularSentence(sentence string) bool {
	n := len(sentence)
	if sentence[0] != sentence[n-1] {
		return false
	}

	for i := 1; i < n; i++ {
		if sentence[i] != ' ' {
			continue
		}
		if sentence[i-1] != sentence[i+1] {
			return false
		}
	}

	return true
}
