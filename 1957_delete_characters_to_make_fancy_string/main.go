package main

func makeFancyString(s string) string {
	count := 1
	result := []byte{s[0]}

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		if count < 3 {
			result = append(result, s[i])
		}
	}

	return string(result)
}
