package string_calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9') || ch == '-'
}

func isDelimiter(ch, delimiter rune) bool {
	if delimiter != 0 {
		return ch == delimiter
	}
	return ch == ',' || ch == '\n'
}

func addNumber(currentSum int, number string, negativeNumbers *[]string) int {
	parsedNumber, _ := strconv.Atoi(number)
	if parsedNumber < 0 {
		*negativeNumbers = append(*negativeNumbers, fmt.Sprintf("%d", parsedNumber))
	}
	sum := currentSum + parsedNumber
	return sum
}

func Sum(expression string) (int, error) {
	result := 0
	number := ""
	var delimiter rune
	if len(expression) > 0 && expression[0] == '\\' {
		delimiter = rune(expression[1])
	}
	var negativeNumbers []string
	for _, ch := range expression {
		if isDigit(ch) {
			number += string(ch)
		}
		if isDelimiter(ch, delimiter) {
			result = addNumber(result, number, &negativeNumbers)
			number = ""
		}
	}
	result = addNumber(result, number, &negativeNumbers)
	if len(negativeNumbers) > 0 {
		return 0, fmt.Errorf("negatives not allowed: [%s]", strings.Join(negativeNumbers, ","))
	}
	return result, nil
}
