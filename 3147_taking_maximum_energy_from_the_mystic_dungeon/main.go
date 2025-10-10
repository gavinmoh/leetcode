package main

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	absorbed := make([]int, n)
	maxEnergy := energy[n-1]

	for i := n - 1; i >= 0; i-- {
		absorbed[i] = energy[i]
		if i+k < n {
			absorbed[i] += absorbed[i+k]
		}

		if absorbed[i] > maxEnergy {
			maxEnergy = absorbed[i]
		}
	}

	return maxEnergy
}
