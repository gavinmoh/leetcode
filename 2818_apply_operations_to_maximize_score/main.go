package main

import (
	"container/heap"
	"math"
)

type Item struct {
	Num   int
	Index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Num > pq[j].Num
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

const MOD = 1_000_000_007

func maximumScore(nums []int, k int) int {
	n := len(nums)
	scores := make([]int, n)
	for i, num := range nums {
		for factor := 2; factor <= int(math.Sqrt(float64(num))); factor++ {
			if num%factor == 0 {
				scores[i]++

				for num%factor == 0 {
					num /= factor
				}
			}
		}

		if num >= 2 {
			scores[i]++
		}
	}

	nextDominant := make([]int, n)
	prevDominant := make([]int, n)
	for i := 0; i < n; i++ {
		nextDominant[i] = n
		prevDominant[i] = -1
	}

	decreasingPrimeScoreStack := []int{}

	for i := 0; i < n; i++ {
		for len(decreasingPrimeScoreStack) > 0 && scores[decreasingPrimeScoreStack[len(decreasingPrimeScoreStack)-1]] < scores[i] {
			m := len(decreasingPrimeScoreStack)
			topIndex := decreasingPrimeScoreStack[m-1]
			decreasingPrimeScoreStack = decreasingPrimeScoreStack[:m-1]
			nextDominant[topIndex] = i
		}

		if len(decreasingPrimeScoreStack) > 0 {
			m := len(decreasingPrimeScoreStack)
			prevDominant[i] = decreasingPrimeScoreStack[m-1]
		}

		decreasingPrimeScoreStack = append(decreasingPrimeScoreStack, i)
	}

	numOfSubarrays := make([]int, n)
	for i := 0; i < n; i++ {
		numOfSubarrays[i] = (nextDominant[i] - i) * (i - prevDominant[i])
	}

	processingQueue := &PriorityQueue{}
	heap.Init(processingQueue)
	for i := 0; i < n; i++ {
		heap.Push(processingQueue, &Item{Num: nums[i], Index: i})
	}

	score := 1

	for k > 0 {
		item := heap.Pop(processingQueue).(*Item)
		operations := min(k, numOfSubarrays[item.Index])
		score = (score * power(item.Num, operations)) % MOD
		k -= operations
	}

	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func power(base int, exponent int) int {
	result := 1

	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % MOD
		}

		base = (base * base) % MOD
		exponent /= 2
	}

	return result
}
