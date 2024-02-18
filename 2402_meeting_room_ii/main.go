package main

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Meeting struct {
	endTime int
	room    int
}

type PriorityQueue []*Meeting

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].endTime == pq[j].endTime {
		return pq[i].room < pq[j].room
	}
	return pq[i].endTime < pq[j].endTime
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	meeting := x.(*Meeting)
	*pq = append(*pq, meeting)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	meeting := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return meeting
}

func (pq PriorityQueue) Top() any {
	return pq[0]
}

func mostBooked(n int, meetings [][]int) int {

	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	rooms := &MinHeap{}
	heap.Init(rooms)
	for i := 0; i < n; i++ {
		heap.Push(rooms, i)
	}

	schedule := &PriorityQueue{}
	heap.Init(schedule)

	usages := make([]int, n)

	for _, meeting := range meetings {
		// free the rooms
		for schedule.Len() > 0 && meeting[0] >= schedule.Top().(*Meeting).endTime {
			ended := heap.Pop(schedule).(*Meeting)
			heap.Push(rooms, ended.room)
		}

		if rooms.Len() > 0 {
			room := heap.Pop(rooms).(int)
			usages[room]++
			newMeeting := &Meeting{room: room, endTime: meeting[1]}
			heap.Push(schedule, newMeeting)
		} else {
			soonest := heap.Pop(schedule).(*Meeting)
			newMeeting := &Meeting{room: soonest.room, endTime: meeting[1]}
			if soonest.endTime > meeting[0] {
				diff := soonest.endTime - meeting[0]
				newMeeting.endTime += diff
			}
			usages[soonest.room]++
			heap.Push(schedule, newMeeting)
		}
	}

	maxRoom, maxUsage := -1, -1
	for room, usage := range usages {
		if usage > maxUsage {
			maxRoom, maxUsage = room, usage
		}
	}

	return maxRoom
}
