package leetcode

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxLength := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if index, ok := m[s[right]]; ok && index >= left {
			left = index + 1
		}
		m[s[right]] = right
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}
	return maxLength
}
