package string_calculator

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SUM      = iota
	MULTIPLY = iota
)

type StringCalculator struct {
	calledCount     int
	negativeNumbers []string
	delimiter       map[string]bool
	operation       int
	result          int
}

func (calc *StringCalculator) parseAndCalculate(number string) {
	parsedNumber, _ := strconv.Atoi(number)
	if parsedNumber < 0 {
		calc.negativeNumbers = append(calc.negativeNumbers, fmt.Sprintf("%d", parsedNumber))
	}
	if parsedNumber >= 1000 {
		return
	}
	switch calc.operation {
	case MULTIPLY:
		calc.result *= parsedNumber
	default:
		calc.result += parsedNumber
	}
}

func (calc *StringCalculator) setDelimiter(expression string) {
	if len(expression) > 0 && expression[0] == '\\' {
		if expression[1] == '[' {
			delimiter := getDelimiterFromExpression(expression)
			if isMultiply(delimiter) {
				calc.operation = MULTIPLY
				calc.result = 1
			}
			calc.delimiter = delimiter
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
	calc.result = calc.operation
}

func (calc *StringCalculator) Calculate(exp string) (int, error) {
	delimiterString, expression := splitExrepressionAndDelimiter(exp)
	calc.setDelimiter(delimiterString)
	defer calc.reset() // similar to finally block in java
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
				calc.parseAndCalculate(number)
				number = ""
			}
			delimiter = append(delimiter, byte(ch))
		}
	}
	calc.parseAndCalculate(number)
	if len(calc.negativeNumbers) > 0 {
		return 0, fmt.Errorf("negatives not allowed: [%s]", strings.Join(calc.negativeNumbers, ","))
	}
	return calc.result, nil
}

func (calc *StringCalculator) GetCalledCount() int {
	return calc.calledCount
}
