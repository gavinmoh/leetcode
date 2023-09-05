package main

var mapping = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	results := []string{}

	for _, digit := range digits {
		// if there is nothing in result, add the first digit
		if len(results) == 0 {
			results = append(results, mapping[string(digit)]...)
		} else {
			// we make a copy of the results
			previousResults := results
			// we reset the results
			results = []string{}
			// for each result, we append the new letters
			for _, result := range previousResults {
				newResults := []string{}
				for _, letter := range mapping[string(digit)] {
					newResults = append(newResults, result+letter)
				}
				// we append the new results to the results
				results = append(results, newResults...)
			}
		}
	}

	return results
}
