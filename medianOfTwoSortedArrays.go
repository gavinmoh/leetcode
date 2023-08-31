// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var merged []int
	var i, j int
	// loop through both arrays as long as there are elements in both
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++ // increment i, since we added it to merged
		} else {
			merged = append(merged, nums2[j])
			j++ // increment j, since we added it to merged
		}
	}
	// it is either nums1 or nums2 that has elements left
	// so we can just append the rest of the elements to merged
	if i < len(nums1) {
		merged = append(merged, nums1[i:]...)
	} else {
		merged = append(merged, nums2[j:]...)
	}

	// if merged is even, we need to get the average of the middle two elements
	if len(merged)%2 == 0 {
		middle := len(merged) / 2
		return float64(merged[middle]+merged[middle-1]) / 2
	} else {
		// if merged is odd, we just need to get the middle element
		return float64(merged[len(merged)/2])
	}
}
