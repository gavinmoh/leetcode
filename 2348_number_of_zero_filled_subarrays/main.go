package main

func zeroFilledSubarray(nums []int) int64 {
	count := int64(0)
	streak := 0

	for _, num := range nums {
		if num == 0 {
			count++
			streak++
		} else {
			count += calculateSubarraysCount(int64(streak))
			streak = 0
		}
	}

	count += calculateSubarraysCount(int64(streak))

	return count
}

func calculateSubarraysCount(streak int64) int64 {
	if streak <= int64(1) {
		return int64(0)
	}

	return (streak - int64(1)) * streak / int64(2)
}
