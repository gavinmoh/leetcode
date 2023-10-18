package main

type Course struct {
	Id          int
	Prequisites []*Course
	Time        int
}

func (c *Course) GetTime(cache []int) int {
	if cache[c.Id] != -1 {
		return cache[c.Id]
	}

	if len(c.Prequisites) == 0 {
		cache[c.Id] = c.Time
		return c.Time
	}

	maxTime := 0
	for _, preq := range c.Prequisites {
		var preqTime int
		if cache[preq.Id] != -1 {
			preqTime = cache[preq.Id]
		} else {
			preqTime = preq.GetTime(cache)
			cache[preq.Id] = preqTime
		}
		if preqTime > maxTime {
			maxTime = preqTime
		}
	}

	total := c.Time + maxTime
	cache[c.Id] = total
	return total
}

func minimumTime(n int, relations [][]int, time []int) int {
	courses := make(map[int]*Course, n+1)
	cache := make([]int, n+1)

	for i := 0; i < n; i++ {
		courses[i+1] = &Course{Id: i + 1, Time: time[i]}
		cache[i+1] = -1
	}

	for _, relation := range relations {
		course := courses[relation[1]]
		course.Prequisites = append(course.Prequisites, courses[relation[0]])
	}

	maxTime := 0
	for i := 1; i <= n; i++ {
		time := courses[i].GetTime(cache)
		if time > maxTime {
			maxTime = time
		}
	}

	return maxTime
}
