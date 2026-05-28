package leetcode

import "strings"

func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	seenDigit := false
	seenDot := false
	seenE := false

	for i, ch := range s {
		switch {
		case ch >= '0' && ch <= '9':
			seenDigit = true

		case ch == '.':
			if seenDot || seenE {
				return false
			}
			seenDot = true

		case ch == 'e' || ch == 'E':
			if seenE || !seenDigit {
				return false
			}
			seenE = true
			seenDigit = false

		case ch == '+' || ch == '-':
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}

		default:
			return false
		}
	}

	return seenDigit
}
