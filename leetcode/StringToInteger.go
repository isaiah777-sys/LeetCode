package leetcode

func myAtoi(s string) int {
	//Note that the algorithm was already given  in the Leetcode question
	i, sign, result := 0, 1, 0
	// Skip leading whitespace
	for i < len(s) && s[i] == ' ' {
		i++
	}
	// Check for optional sign
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	// Convert digits to integer
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i]-'0')
		// Check for overflow
		if result > 1<<31-1 {
			if sign == 1 {
				return 1<<31 - 1
			}
			return -1 << 31
		}
		i++
	}
	return sign * result
}
