package main

func canChange(start string, target string) bool {
	startIndex, targetIndex := 0, 0
	startLength, targetLength := len(start), len(target)

	if startLength != targetLength {
		return false
	}

	for startIndex < startLength || targetIndex < targetLength {
		for startIndex < startLength && start[startIndex] == '_' {
			startIndex++
		}

		for targetIndex < targetLength && target[targetIndex] == '_' {
			targetIndex++
		}

		if startIndex == startLength || targetIndex == startLength {
			return startIndex == startLength && targetIndex == startLength
		}

		if start[startIndex] != target[targetIndex] {
			return false
		}

		if start[startIndex] == 'L' && startIndex < targetIndex {
			return false
		} else if start[startIndex] == 'R' && startIndex > targetIndex {
			return false
		}

		startIndex++
		targetIndex++
	}

	return true
}
