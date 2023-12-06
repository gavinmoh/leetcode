package main

func totalMoney(n int) int {
	// arithmetic sum
	// sum = k * (F + L) / 2
	// k is number of full week
	// F is first element in the sequence
	// L is the last element in the sequence

	k := n / 7
	// 1 + 2 + 3 + ... + 7
	F := 28
	// extra 1 coin for each monday, so each week will have extra 7 coins
	L := 28 + (k-1)*7
	sum := k * (F + L) / 2

	remainder := n % 7
	finalWeek := 0
	monday := 1 + k
	for i := 0; i < remainder; i++ {
		finalWeek += monday + i
	}

	return sum + finalWeek

}
