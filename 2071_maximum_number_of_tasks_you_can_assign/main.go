package main

import (
	"container/list"
	"sort"
)

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	left, right := 0, min(len(tasks), len(workers))

	sort.Ints(tasks)
	sort.Ints(workers)

	answer := 0
	for left <= right {
		mid := (left + right) / 2
		currentWorker := len(workers) - 1
		pillsLeft := pills
		canAssignAll := true
		workerList := list.New()

		for currentTask := mid - 1; currentTask >= 0; currentTask-- {
			for currentWorker >= len(workers)-mid && workers[currentWorker]+strength >= tasks[currentTask] {
				workerList.PushFront(workers[currentWorker])
				currentWorker--
			}

			if workerList.Len() == 0 {
				canAssignAll = false
				break
			}

			if workerList.Back().Value.(int) >= tasks[currentTask] {
				workerList.Remove(workerList.Back())
			} else {
				if pillsLeft == 0 {
					canAssignAll = false
					break
				}
				pillsLeft--
				workerList.Remove(workerList.Front())
			}
		}

		if canAssignAll {
			answer = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return answer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
