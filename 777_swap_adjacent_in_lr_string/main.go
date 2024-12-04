package main

func canTransform(start string, result string) bool {
	startIndex, resultIndex := 0, 0
	startLength, resultLength := len(start), len(result)

	if startLength != resultLength {
		return false
	}

	for startIndex < startLength || resultIndex < resultLength {
		for startIndex < startLength && start[startIndex] == 'X' {
			startIndex++
		}

		for resultIndex < resultLength && result[resultIndex] == 'X' {
			resultIndex++
		}

		if startIndex == startLength || resultIndex == startLength {
			return startIndex == startLength && resultIndex == startLength
		}

		if start[startIndex] != result[resultIndex] {
			return false
		}

		if start[startIndex] == 'L' && startIndex < resultIndex {
			return false
		} else if start[startIndex] == 'R' && startIndex > resultIndex {
			return false
		}

		startIndex++
		resultIndex++
	}

	return true
}
