package main

func subarraysDivByK(nums []int, k int) int {
	mods := make([]int, k)
	prefixMod := 0
	mods[0] = 1
	result := 0

	for _, num := range nums {
		prefixMod = (prefixMod + num) % k
		if prefixMod < 0 {
			prefixMod = (prefixMod + k) % k
		}

		result += mods[prefixMod]
		mods[prefixMod]++
	}

	return result
}
