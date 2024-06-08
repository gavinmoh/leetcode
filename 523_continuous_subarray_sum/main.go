package main

func checkSubarraySum(nums []int, k int) bool {
	mods := make(map[int]int) // storing prefixMod, key is result, value is index
	prefixMod := 0
	mods[0] = -1

	for i, num := range nums {
		prefixMod = (prefixMod + num) % k

		if idx, ok := mods[prefixMod]; ok {
			if i-idx > 1 {
				return true
			}
		} else {
			mods[prefixMod] = i
		}
	}

	return false
}
