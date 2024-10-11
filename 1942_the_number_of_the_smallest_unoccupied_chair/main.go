package main

import (
	"container/heap"
	"sort"
)

type Friend struct {
	Index   int
	Arrival int
	Leaving int
}

// An Item is something we manage in a priority queue.
type Item struct {
	Index    int
	Priority int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func smallestChair(times [][]int, targetFriend int) int {
	friends := make([]*Friend, len(times))
	for friend, time := range times {
		friends[friend] = &Friend{
			Index:   friend,
			Arrival: time[0],
			Leaving: time[1],
		}
	}
	sort.SliceStable(friends, func(i, j int) bool {
		return friends[i].Arrival < friends[j].Arrival
	})

	chairIndex := 0
	empty := &PriorityQueue{}
	occupied := &PriorityQueue{}
	heap.Init(empty)
	heap.Init(occupied)

	for _, friend := range friends {
		for occupied.Len() > 0 {
			chair := heap.Pop(occupied).(*Item)
			if chair.Priority <= friend.Arrival {
				// put the chair into empty queue
				heap.Push(empty, &Item{
					Index:    chair.Index,
					Priority: chair.Index,
				})
			} else {
				// put back the chair
				heap.Push(occupied, chair)
				break
			}
		}

		if empty.Len() > 0 {
			// grab the first empty chair
			chair := heap.Pop(empty).(*Item)
			heap.Push(occupied, &Item{
				Index:    chair.Index,
				Priority: friend.Leaving,
			})
			if friend.Index == targetFriend {
				return chair.Index
			}
		} else {
			heap.Push(occupied, &Item{
				Index:    chairIndex,
				Priority: friend.Leaving,
			})

			if friend.Index == targetFriend {
				return chairIndex
			}

			chairIndex++
		}
	}

	return -1
}
