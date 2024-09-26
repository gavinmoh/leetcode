package main

type MyCalendar struct {
	Intervals [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{Intervals: [][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	// check if we can insert in front
	if len(this.Intervals) == 0 || this.Intervals[0][0] >= end {
		this.Intervals = append([][]int{{start, end}}, this.Intervals...)
		return true
	}

	for i := 0; i < len(this.Intervals)-1; i++ {
		if start >= this.Intervals[i][1] && end <= this.Intervals[i+1][0] {
			left := this.Intervals[:i+1]
			right := make([][]int, len(this.Intervals[i+1:]))
			copy(right, this.Intervals[i+1:])
			this.Intervals = append(left, []int{start, end})
			this.Intervals = append(this.Intervals, right...)
			return true
		}
	}

	// check if we can insert the back
	if this.Intervals[len(this.Intervals)-1][1] <= start {
		this.Intervals = append(this.Intervals, []int{start, end})
		return true
	}

	return false
}

/**
* Your MyCalendar object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Book(start,end);
 */
