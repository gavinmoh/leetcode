package main

import (
	"fmt"
	"strconv"
)

func findDifferentBinaryString(nums []string) string {
	numsMap := make(map[int]bool)
	binaryLength := len(nums[0])
	formatStr := "%0" + fmt.Sprint(binaryLength) + "b"

	maxNum := 0
	for _, num := range nums {
		numInt, _ := strconv.ParseInt(num, 2, 16)
		maxNum = max(maxNum, int(numInt))
		numsMap[int(numInt)] = true
	}

	for i := 0; i <= maxNum; i++ {
		if !numsMap[i] {
			return fmt.Sprintf(formatStr, i)
		}
	}

	return fmt.Sprintf(formatStr, maxNum+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
