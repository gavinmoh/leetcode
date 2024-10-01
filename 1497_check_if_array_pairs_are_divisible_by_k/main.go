package main

func canArrange(arr []int, k int) bool {
	freq := make(map[int]int)
	for _, num := range arr {
		remainder := ((num % k) + k) % k
		freq[remainder] += 1
	}

	for _, num := range arr {
		remainder := ((num % k) + k) % k
		freq1 := freq[remainder]
		freq2 := freq[k-remainder]

		if remainder == 0 {
			if freq1%2 == 1 {
				return false
			}
		} else if freq1 != freq2 {
			return false
		}
	}

	return true
}
