package string_calculator

import (
	"fmt"
	"strconv"
)

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isDelimiter(ch, delimiter rune) bool {
	if delimiter != 0 {
		return ch == delimiter
	}
	return ch == ',' || ch == '\n'
}

func Sum(expression string) (int, error) {
	result := 0
	number := ""
	isNegative := false
	var delimiter rune
	if len(expression) > 0 && expression[0] == '\\' {
		delimiter = rune(expression[1])
	}
	for _, ch := range expression {
		if isDigit(ch) {
			number += string(ch)
		}
		if isDelimiter(ch, delimiter) {
			parsedNumber, _ := strconv.Atoi(number)
			if isNegative {
				return 0, fmt.Errorf("negatives not allowed: [%d]", (-1 * parsedNumber))
			}
			result += parsedNumber
			number = ""
		}
		if ch == '-' {
			isNegative = true
		}
	}
	parsedNumber, _ := strconv.Atoi(number)
	if isNegative {
		return 0, fmt.Errorf("negatives not allowed: [%d]", (-1 * parsedNumber))
	}
	result += parsedNumber
	return result, nil
}
