package main

import "fmt"

var (
	HOURS_MAP = [][]string{
		{"0"},
		{"8", "4", "2", "1"},
		{"10", "9", "6", "5", "3"},
		{"11", "7"},
	}

	MINUTES_MAP = [][]string{
		{"00"},                               // 1
		{"32", "16", "08", "04", "02", "01"}, // 6
		{"48", "40", "36", "34", "33", "24", "20", "18", "17", "12", "10", "09", "06", "05", "03"},                               // 15
		{"56", "52", "50", "49", "44", "42", "41", "38", "37", "35", "28", "26", "25", "22", "21", "19", "14", "13", "11", "07"}, // 20
		{"58", "57", "54", "53", "51", "46", "45", "43", "39", "30", "29", "27", "23", "15"},                                     // 14
		{"59", "55", "47", "31"}, // 4
	}
)

func readBinaryWatch(turnedOn int) []string {
	result := []string{}

	for i, hours := range HOURS_MAP {
		for j, minutes := range MINUTES_MAP {
			if i+j != turnedOn {
				continue
			}

			for _, hour := range hours {
				for _, minute := range minutes {
					result = append(result, fmt.Sprintf("%s:%s", hour, minute))
				}
			}
		}
	}

	return result
}
