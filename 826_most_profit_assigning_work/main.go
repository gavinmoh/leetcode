package main

import "sort"

type Job struct {
	Difficulty int
	Profit     int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	jobs := make([]Job, len(difficulty))
	for i := 0; i < len(difficulty); i++ {
		jobs[i] = Job{Difficulty: difficulty[i], Profit: profit[i]}
	}

	// sort job by profit and difficulty descending
	sort.SliceStable(jobs, func(i, j int) bool {
		if jobs[i].Profit == jobs[j].Profit {
			return jobs[i].Difficulty > jobs[j].Difficulty
		}
		return jobs[i].Profit > jobs[j].Profit
	})

	// sort worker ability descending
	sort.SliceStable(worker, func(i, j int) bool {
		return worker[i] > worker[j]
	})

	sum := 0
	workerIdx, jobIdx := 0, 0
	// looping through worker, if job is more difficult, then skip to next job
	for workerIdx < len(worker) && jobIdx < len(jobs) {
		if jobs[jobIdx].Difficulty > worker[workerIdx] {
			jobIdx++
		} else {
			sum += jobs[jobIdx].Profit
			workerIdx++
		}
	}

	return sum
}
