package string_calculator

import "strconv"

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isSeparator(ch rune) bool {
	return ch == ',' || ch == '\n'
}

func Sum(numbers string) int {
	result := 0
	number := ""
	for _, ch := range numbers {
		if isDigit(ch) {
			number += string(ch)
		}
		if isSeparator(ch) {
			parsedNumber, _ := strconv.Atoi(number)
			result += parsedNumber
			number = ""
		}
	}
	parsedNumber, _ := strconv.Atoi(number)
	result += parsedNumber
	return result
}
