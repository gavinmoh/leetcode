package main

func parseBoolExpr(expression string) bool {
	stack := []rune{}
	for _, char := range expression {
		switch char {
		case '(', ',':
			continue
		case ')':
			// pop until operator
			subexpr := []rune{}
			var result bool
			for len(stack) > 0 {
				c := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if isOperator(c) {
					result = parse(c, subexpr)
					break
				} else {
					subexpr = append([]rune{c}, subexpr...)
				}
			}
			stack = append(stack, toChar(result))
		default:
			stack = append(stack, char)
		}
	}

	return toBool(stack[0])
}

func toBool(c rune) bool {
	return c == 't'
}

func toChar(b bool) rune {
	if b {
		return 't'
	}
	return 'f'
}

func isOperator(c rune) bool {
	switch c {
	case '!', '&', '|':
		return true
	default:
		return false
	}
}

func parse(operator rune, expr []rune) bool {
	switch operator {
	case '!':
		return !toBool(expr[0])
	case '&':
		result := toBool(expr[0])
		for i := 1; i < len(expr); i++ {
			result = result && toBool(expr[i])
		}
		return result
	case '|':
		result := toBool(expr[0])
		for i := 1; i < len(expr); i++ {
			result = result || toBool(expr[i])
		}
		return result
	}

	return false
}
