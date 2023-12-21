package string_calculator

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9') || ch == '-'
}
