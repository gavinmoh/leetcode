package main

import "strings"

var numbers = map[int]string{
	0:             "",
	1:             "One",
	2:             "Two",
	3:             "Three",
	4:             "Four",
	5:             "Five",
	6:             "Six",
	7:             "Seven",
	8:             "Eight",
	9:             "Nine",
	10:            "Ten",
	11:            "Eleven",
	12:            "Twelve",
	13:            "Thirteen",
	14:            "Fourteen",
	15:            "Fifteen",
	16:            "Sixteen",
	17:            "Seventeen",
	18:            "Eighteen",
	19:            "Nineteen",
	20:            "Twenty",
	30:            "Thirty",
	40:            "Forty",
	50:            "Fifty",
	60:            "Sixty",
	70:            "Seventy",
	80:            "Eighty",
	90:            "Ninety",
	100:           "Hundred",
	1000:          "Thousand",
	1_000_000:     "Million",
	1_000_000_000: "Billion",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := []string{}

	billion := num / 1_000_000_000
	if billion > 0 {
		result = append(result, numbers[billion])
		result = append(result, numbers[1_000_000_000])
		num %= 1_000_000_000
	}

	million := num / 1_000_000
	if million > 0 {
		result = append(result, helper(million%1_000_000))
		result = append(result, numbers[1_000_000])
		num %= 1_000_000
	}

	thousand := num / 1_000
	if thousand > 0 {
		result = append(result, helper(thousand%1_000))
		result = append(result, numbers[1_000])
		num %= 1_000
	}

	if num > 0 {
		result = append(result, helper(num))
	}

	return strings.Join(result, " ")
}

func helper(num int) string {
	result := []string{}
	hundredth := num / 100
	if hundredth > 0 {
		result = append(result, numbers[hundredth])
		result = append(result, numbers[100])
	}
	tenth := num % 100

	if tenth > 20 {
		result = append(result, numbers[(tenth/10)*10])
		if (tenth % 10) > 0 {
			result = append(result, numbers[tenth%10])
		}
	} else if tenth > 0 {
		result = append(result, numbers[tenth])
	}

	return strings.Join(result, " ")
}
