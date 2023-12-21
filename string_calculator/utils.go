package string_calculator

import (
	"strings"
)

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9') || ch == '-'
}

func extractExpression(expression string) string {
	if len(expression) == 0 || expression[0] != '\\' {
		return expression
	}
	ind := 0
	for expression[ind] < '0' || expression[ind] > '9' {
		ind++
	}
	return strings.Trim(expression[ind:], " ")
}
