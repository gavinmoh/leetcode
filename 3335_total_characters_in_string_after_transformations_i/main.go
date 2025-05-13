package main

const MOD = 1e9 + 7

func lengthAfterTransformations(s string, t int) int {
	freq := make([]int, 26)
	for _, c := range s {
		freq[int(c-'a')]++
	}

	for i := t; i > 0; i-- {
		newFreq := make([]int, 26)

		for c := 0; c < 25; c++ {
			newFreq[c+1] = freq[c]
		}

		if freq[25] > 0 {
			newFreq[0] = (newFreq[0] + freq[25]) % MOD
			newFreq[1] = (newFreq[1] + freq[25]) % MOD
		}

		freq = newFreq
	}

	result := 0
	for i := 0; i < 26; i++ {
		result = (result + freq[i]) % MOD
	}

	return result
}
