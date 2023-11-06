package main

import (
	"sort"
)

type SeatManager struct {
	seats         map[int]bool
	priorityQueue []int
	nextSeat      int
}

func Constructor(n int) SeatManager {
	seats := make(map[int]bool, n)
	for i := 1; i <= n; i++ {
		seats[i] = false
	}

	return SeatManager{seats: seats, priorityQueue: []int{}, nextSeat: 1}
}

func (this *SeatManager) Reserve() int {
	if len(this.priorityQueue) == 0 {
		seatNumber := this.nextSeat
		this.nextSeat++
		this.seats[seatNumber] = true
		return seatNumber
	}
	seatNumber := this.priorityQueue[0]
	this.priorityQueue = this.priorityQueue[1:]
	this.seats[seatNumber] = true
	return seatNumber
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.seats[seatNumber] = false
	this.priorityQueue = append(this.priorityQueue, seatNumber)
	sort.Ints(this.priorityQueue)
}
