package main

type MountainArray struct {
	data     []int
	apiCalls int `default:"0"`
}

func (this *MountainArray) get(index int) int {
	if this.apiCalls > 100 {
		panic("API calls limit exceeded")
	}
	this.apiCalls++
	return this.data[index]
}

func (this *MountainArray) length() int {
	return len(this.data)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()

	// find the peak
	left := 1
	right := length - 2
	var mid int

	for left <= right {
		mid = left + (right-left)/2
		midValue := mountainArr.get(mid)
		midLeftValue := mountainArr.get(mid - 1)
		midRightValue := mountainArr.get(mid + 1)

		// found the peak
		if midValue > midLeftValue && midValue > midRightValue {
			break
		}

		// ascending
		if midValue > midLeftValue && midValue < midRightValue {
			left = mid + 1
		} else { // descending
			right = mid - 1
		}
	}

	// find the target in the left side
	peak := mid
	left = 0
	right = peak
	for left <= right {
		mid = left + (right-left)/2
		midValue := mountainArr.get(mid)
		if midValue == target {
			return mid
		}

		if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// find the target in the right side
	left = peak
	right = length - 1
	for left <= right {
		mid = left + (right-left)/2
		midValue := mountainArr.get(mid)
		if midValue == target {
			return mid
		}

		if midValue > target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
