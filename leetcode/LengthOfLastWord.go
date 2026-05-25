package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}

//Another approach : using two pointers technique to find the length of the last word in a string. It iterates through the string from the end, skipping any trailing spaces, and then counts the characters of the last word until it encounters a space or reaches the beginning of the string.
func lengthOfLastWord2(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length > 0 {
				break
			} else {
				continue
			}
		}
		length++
	}
	return length
}
