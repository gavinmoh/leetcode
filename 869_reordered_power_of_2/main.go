package main

func reorderedPowerOf2(n int) bool {
	if n == 1 {
		return true
	}

	freq, numDigits := numToFreq(n)
	minNum, maxNum := 1, 9
	for i := 1; i < numDigits; i++ {
		minNum *= 10
		maxNum = (maxNum * 10) + 9
	}

	powOfTwo := 1
	for powOfTwo <= maxNum {
		powOfTwo *= 2
		if powOfTwo < minNum {
			continue
		}

		freq2, _ := numToFreq(powOfTwo)
		matched := true
		for i := 0; i < 10; i++ {
			if freq[i] != freq2[i] {
				matched = false
				break
			}
		}

		if matched {
			return true
		}
	}

	return false
}

func numToFreq(n int) ([]int, int) {
	freq := make([]int, 10)
	numDigits := 0

	for num := n; num > 0; num = num / 10 {
		digit := num % 10
		freq[digit]++
		numDigits++
		n /= 10
	}

	return freq, numDigits
}
