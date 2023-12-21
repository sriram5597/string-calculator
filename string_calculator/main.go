package string_calculator

import (
	"fmt"
	"strconv"
	"strings"
)

type StringCalculator struct {
	calledCount     int
	negativeNumbers []string
	delimiter       map[string]bool
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
		if expression[1] == '[' {
			calc.delimiter = getDelimiterFromExpression(expression)
		} else {
			calc.delimiter = map[string]bool{
				string(expression[1]): true,
			}
		}
	} else {
		calc.delimiter = map[string]bool{
			"\n": true,
			",":  true,
		}
	}
}

func (calc *StringCalculator) isValidDelimiter(delimiter string) bool {
	_, ok := calc.delimiter[delimiter]
	return delimiter == "" || ok
}

func (calc *StringCalculator) reset() {
	calc.calledCount++
	calc.negativeNumbers = []string{}
	calc.delimiter = map[string]bool{}
}

func (calc *StringCalculator) Sum(exp string) (int, error) {
	delimiterString, expression := splitExrepressionAndDelimiter(exp)
	calc.setDelimiter(delimiterString)
	defer calc.reset() // similar to finally block in java
	result := 0
	number := ""
	delimiter := []byte{}
	for _, ch := range expression {
		if isDigit(ch) {
			if !calc.isValidDelimiter(string(delimiter)) {
				return 0, fmt.Errorf("invalid expression")
			}
			if len(number) == 0 {
				delimiter = []byte{}
			}
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
