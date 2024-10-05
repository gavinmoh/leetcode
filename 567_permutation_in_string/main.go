package main

func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)

	s1Count := make([]int, 26)
	for _, letter := range s1 {
		s1Count[letter-'a'] += 1
	}

	s2Count := make([]int, 26)
	right := n
	for left := 0; left <= len(s2)-n; left++ {
		if left == 0 {
			for i := 0; i < right; i++ {
				letter := s2[i]
				s2Count[letter-'a'] += 1
			}
		} else {
			leftLetter := s2[left-1]
			rightLetter := s2[left+n-1]
			s2Count[leftLetter-'a'] -= 1
			s2Count[rightLetter-'a'] += 1
		}

		isSubstring := true
		for i := 0; i < 26; i++ {
			if s1Count[i] != s2Count[i] {
				isSubstring = false
				break
			}
		}
		if isSubstring {
			return true
		}
	}

	return false
}
