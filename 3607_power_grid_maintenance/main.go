package main

type DSU struct {
	parent []int
}

func NewDSU(size int) *DSU {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}

	return &DSU{parent}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] == x {
		return x
	}

	d.parent[x] = d.Find(d.parent[x])
	return d.parent[x]
}

func (d *DSU) Join(u, v int) {
	d.parent[d.Find(v)] = d.Find(u)
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	dsu := NewDSU(c + 1)

	for _, connection := range connections {
		u, v := connection[0], connection[1]
		dsu.Join(u, v)
	}

	online := make([]bool, c+1)
	for i := range online {
		online[i] = true
	}

	offlineCounts := make([]int, c+1)
	for _, query := range queries {
		op, x := query[0], query[1]
		if op == 2 {
			online[x] = false
			offlineCounts[x]++
		}
	}

	minimumOnlineStations := map[int]int{}
	for i := 1; i <= c; i++ {
		root := dsu.Find(i)
		if _, exists := minimumOnlineStations[root]; !exists {
			minimumOnlineStations[root] = -1
		}
		station := minimumOnlineStations[root]
		if online[i] {
			if station == -1 || station > i {
				minimumOnlineStations[root] = i
			}
		}
	}

	result := []int{}
	for i := len(queries) - 1; i >= 0; i-- {
		query := queries[i]
		op, x := query[0], query[1]
		root := dsu.Find(x)
		station := minimumOnlineStations[root]

		if op == 1 {
			if online[x] {
				result = append(result, x)
			} else {
				result = append(result, station)
			}
		}

		if op == 2 {
			if offlineCounts[x] > 1 {
				offlineCounts[x]--
			} else {
				online[x] = true
				if station == -1 || station > x {
					minimumOnlineStations[root] = x
				}
			}
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
