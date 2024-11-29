package main

func checkIfExist(arr []int) bool {
	nums := map[int]bool{}
	for _, num := range arr {
		if _, ok := nums[num*2]; ok {
			return true
		}

		if num%2 == 0 {
			if _, ok := nums[num/2]; ok {
				return true
			}
		}

		nums[num] = true
	}

	return false
}
