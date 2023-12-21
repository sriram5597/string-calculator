package string_calculator

import (
	"strings"
)

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9') || ch == '-'
}

func splitExrepressionAndDelimiter(expression string) (string, string) {
	if len(expression) == 0 || expression[0] != '\\' {
		return "", expression
	}
	ind := 0
	for expression[ind] < '0' || expression[ind] > '9' {
		ind++
	}
	return expression[:ind], strings.Trim(expression[ind:], " ")
}

func getDelimiterFromExpression(expression string) map[string]bool {
	ind := 2
	delimiter := []byte{}
	delimiterMap := map[string]bool{}
	for expression[ind] != '\n' {
		if expression[ind] == ']' {
			delimiterMap[string(delimiter)] = true
			delimiter = []byte{}
		} else if expression[ind] != '[' {
			delimiter = append(delimiter, expression[ind])
		}
		ind++
	}
	return delimiterMap
}
