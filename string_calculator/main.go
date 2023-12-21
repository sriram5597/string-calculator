package string_calculator

import (
	"fmt"
	"strconv"
	"strings"
)

type StringCalculator struct {
	calledCount     int
	negativeNumbers []string
	delimiter       rune
}

func (calc *StringCalculator) parseAndAdd(currentSum int, number string) int {
	parsedNumber, _ := strconv.Atoi(number)
	if parsedNumber < 0 {
		calc.negativeNumbers = append(calc.negativeNumbers, fmt.Sprintf("%d", parsedNumber))
	}
	if parsedNumber >= 1000 {
		return currentSum
	}
	sum := currentSum + parsedNumber
	return sum
}

func (calc *StringCalculator) setDelimiter(expression string) {
	if len(expression) > 0 && expression[0] == '\\' {
		calc.delimiter = rune(expression[1])
	}
}

func (calc *StringCalculator) isDelimiter(ch rune) bool {
	if calc.delimiter != 0 {
		return ch == calc.delimiter
	}
	return ch == ',' || ch == '\n'
}

func (calc *StringCalculator) reset() {
	calc.negativeNumbers = []string{}
	calc.delimiter = 0
}

func (calc *StringCalculator) Sum(expression string) (int, error) {
	calc.reset()
	calc.setDelimiter(expression)
	result := 0
	number := ""
	for _, ch := range expression {
		if isDigit(ch) {
			number += string(ch)
		}
		if calc.isDelimiter(ch) {
			result = calc.parseAndAdd(result, number)
			number = ""
		}
	}
	result = calc.parseAndAdd(result, number)
	calc.calledCount++
	if len(calc.negativeNumbers) > 0 {
		return 0, fmt.Errorf("negatives not allowed: [%s]", strings.Join(calc.negativeNumbers, ","))
	}
	return result, nil
}

func (calc *StringCalculator) GetCalledCount() int {
	return calc.calledCount
}
