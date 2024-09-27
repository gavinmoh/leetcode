package main

type MyCalendarTwo struct {
	NonOverlapping [][]int
	Overlapping    [][]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		NonOverlapping: [][]int{},
		Overlapping:    [][]int{},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, event := range this.Overlapping {
		if end > event[0] && event[1] > start {
			return false
		}
	}

	for _, event := range this.NonOverlapping {
		if end > event[0] && event[1] > start {
			this.Overlapping = append(this.Overlapping, []int{max(start, event[0]), min(end, event[1])})
		}
	}
	this.NonOverlapping = append(this.NonOverlapping, []int{start, end})

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
* Your MyCalendarTwo object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Book(start,end);
 */
