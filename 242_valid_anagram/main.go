package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_map := make(map[rune]int, len(s))
	t_map := make(map[rune]int, len(t))

	for _, char := range s {
		s_map[char]++
	}

	for _, char := range t {
		t_map[char]++
	}

	for char, s_count := range s_map {
		if t_count, ok := t_map[char]; !ok || t_count != s_count {
			return false
		}
	}
	return true
}
