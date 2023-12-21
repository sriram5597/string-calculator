package string_calculator

import (
	"fmt"
	"strconv"
	"strings"
)

type StringCalculator struct {
	calledCount     int
	negativeNumbers []string
	delimiter       string
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
	delimiter := []byte{}
	if len(expression) > 0 && expression[0] == '\\' {
		if expression[1] == '[' {
			ind := 2
			for expression[ind] != ']' {
				delimiter = append(delimiter, expression[ind])
				ind++
			}
		} else {
			delimiter = append(delimiter, expression[1])
		}
	}
	calc.delimiter = string(delimiter)
}

func (calc *StringCalculator) isValidDelimiter(delimiter string) bool {
	return delimiter == "" || delimiter == "," || delimiter == "\n" || delimiter == calc.delimiter
}

func (calc *StringCalculator) reset() {
	calc.negativeNumbers = []string{}
	calc.delimiter = ""
}

func (calc *StringCalculator) Sum(expression string) (int, error) {
	calc.setDelimiter(expression)
	defer func() {
		calc.calledCount++
		calc.reset()
	}()
	result := 0
	number := ""
	delimiter := []byte{}
	extractedExpression := extractExpression(expression)
	for _, ch := range extractedExpression {
		if isDigit(ch) {
			if !calc.isValidDelimiter(string(delimiter)) {
				return 0, fmt.Errorf("invalid expression")
			}
			delimiter = []byte{}
			number += string(ch)
		} else {
			if len(delimiter) == 0 {
				result = calc.parseAndAdd(result, number)
				number = ""
			}
			delimiter = append(delimiter, byte(ch))
		}
	}
	result = calc.parseAndAdd(result, number)
	if len(calc.negativeNumbers) > 0 {
		return 0, fmt.Errorf("negatives not allowed: [%s]", strings.Join(calc.negativeNumbers, ","))
	}
	return result, nil
}

func (calc *StringCalculator) GetCalledCount() int {
	return calc.calledCount
}
