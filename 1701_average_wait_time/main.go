package main

func averageWaitingTime(customers [][]int) float64 {
	currentTime := 0
	sumWaitTime := 0
	for i := 0; i < len(customers); i++ {
		if customers[i][0] > currentTime {
			currentTime = customers[i][0]
		}
		completeTime := currentTime + customers[i][1]
		sumWaitTime += (completeTime - customers[i][0])
		currentTime = completeTime
	}

	return float64(sumWaitTime) / float64(len(customers))
}
