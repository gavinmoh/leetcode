package main

func kthCharacter(k int) byte {
	word := []rune("a")
	for len(word) < k {
		for _, c := range word {
			word = append(word, (c+1)%'z')
		}
	}

	return byte(word[k-1])
}
