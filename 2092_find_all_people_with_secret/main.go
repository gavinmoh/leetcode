package main

import (
	"container/list"
	"math"
)

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	graph := make(map[int][][]int)
	for _, meeting := range meetings {
		x, y, t := meeting[0], meeting[1], meeting[2]
		graph[x] = append(graph[x], []int{t, y})
		graph[y] = append(graph[y], []int{t, x})
	}

	earliest := make([]int, n)
	for i := range earliest {
		earliest[i] = math.MaxInt
	}
	earliest[0] = 0
	earliest[firstPerson] = 0

	q := list.New()
	q.PushBack([]int{0, 0})
	q.PushBack([]int{firstPerson, 0})

	for q.Len() > 0 {
		elem := q.Front()
		q.Remove(elem)
		node := elem.Value.([]int)
		person, time := node[0], node[1]
		for _, pair := range graph[person] {
			t := pair[0]
			nextPerson := pair[1]

			if t >= time && earliest[nextPerson] > t {
				earliest[nextPerson] = t
				q.PushBack([]int{nextPerson, t})
			}
		}
	}

	result := []int{}
	for i := 0; i < n; i++ {
		if earliest[i] != math.MaxInt {
			result = append(result, i)
		}
	}

	return result
}
