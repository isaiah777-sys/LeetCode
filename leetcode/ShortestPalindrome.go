package leetcode

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	reversed := reverseString(s)
	combined := s + "#" + reversed
	lps := computeLPS(combined)
	toAdd := reversed[:len(s)-lps[len(combined)-1]]
	return toAdd + s
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func computeLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	for i := 1; i < len(pattern); {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}

	}
	return lps
}
