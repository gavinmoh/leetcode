package main

import (
	"fmt"
	"strconv"
)

func fractionAddition(expression string) string {
	numerators := []int{}
	denominators := []int{}
	operators := []rune{}

	substring := []rune{}
	for i, char := range expression {
		if i != 0 && (char == '-' || char == '+') {
			numerator, denominator := parseFraction(substring)
			numerators = append(numerators, numerator)
			denominators = append(denominators, denominator)
			substring = []rune{}
			operators = append(operators, char)
		} else {
			substring = append(substring, char)
		}
	}
	numerator, denominator := parseFraction(substring)
	numerators = append(numerators, numerator)
	denominators = append(denominators, denominator)

	commonDenominator := 1
	for i := 0; i < len(denominators); i++ {
		commonDenominator = lcm(commonDenominator, denominators[i])
	}

	for i := 0; i < len(numerators); i++ {
		if denominators[i] == commonDenominator {
			continue
		}
		numerators[i] *= commonDenominator / denominators[i]
	}

	stack := []int{numerators[0]}
	i := 1
	for _, operator := range operators {
		for len(stack) < 2 {
			stack = append(stack, numerators[i])
			i++
		}

		if operator == '+' {
			stack[0] += stack[1]
		} else {
			stack[0] -= stack[1]
		}
		stack = stack[:1]
	}
	if stack[0] == 0 {
		return "0/1"
	}

	divisor := gcd(commonDenominator, abs(stack[0]))
	return fmt.Sprintf("%d/%d", stack[0]/divisor, commonDenominator/divisor)
}

func parseFraction(s []rune) (int, int) {
	numerator := []rune{}
	denominator := []rune{}
	isNumerator := true
	for _, char := range s {
		if char == '/' {
			isNumerator = false
			continue
		}
		if isNumerator {
			numerator = append(numerator, char)
		} else {
			denominator = append(denominator, char)
		}
	}
	numeratorInt, _ := strconv.Atoi(string(numerator))
	denominatorInt, _ := strconv.Atoi(string(denominator))

	return numeratorInt, denominatorInt
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
