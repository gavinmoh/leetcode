package main

func numOfSubarrays(arr []int) int {
	mod := 1_000_000_007
	n := len(arr)

	for i, num := range arr {
		arr[i] = num % 2
	}

	dpEven, dpOdd := make([]int, n), make([]int, n)
	if arr[n-1] == 0 {
		dpEven[n-1] = 1
	} else {
		dpOdd[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if arr[i] == 1 {
			dpOdd[i] += (1 + dpEven[i+1])
			dpEven[i] = dpOdd[i+1]
		} else {
			dpEven[i] += (1 + dpEven[i+1])
			dpOdd[i] = dpOdd[i+1]
		}
	}

	count := 0
	for _, num := range dpOdd {
		count += num
		count %= mod
	}

	return count
}
