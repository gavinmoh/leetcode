package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	criticalPoints := []int{}
	i := 1
	previous, current := head, head.Next

	for current.Next != nil {
		isLocalMaxima := current.Val > previous.Val && current.Val > current.Next.Val
		isLocalMinima := current.Val < previous.Val && current.Val < current.Next.Val
		if isLocalMaxima || isLocalMinima {
			criticalPoints = append(criticalPoints, i)
		}

		previous, current = current, current.Next
		i++
	}

	if len(criticalPoints) < 2 {
		return []int{-1, -1}
	}

	maxDistance := criticalPoints[len(criticalPoints)-1] - criticalPoints[0]
	minDistance := criticalPoints[1] - criticalPoints[0]
	for j := 1; j < len(criticalPoints)-1; j++ {
		currentDistance := criticalPoints[j+1] - criticalPoints[j]
		if currentDistance < minDistance {
			minDistance = currentDistance
		}
	}

	return []int{minDistance, maxDistance}
}
