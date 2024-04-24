package main

func tribonacci(n int) int {
	f := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		f = append(f, f[i-3]+f[i-2]+f[i-1])
	}

	return f[n]
}
