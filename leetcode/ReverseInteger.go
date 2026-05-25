package leetcode

import "strconv"

func reverse(x int) int {
	tmp := x
	if x < 0 {
		tmp = -x
	}
	result := 0
	for tmp > 0 {
		result = result*10 + tmp%10
		tmp /= 10
	}
	if result > 1<<31-1 || result < -1<<31 {
		return 0
	}
	if x < 0 {
		return -result
	}
	return result
}

// Another alternative solution using string manipulation technique to reverse the integer. It converts the integer to a string, reverses the string, and then converts it back to an integer. It also handles the negative sign and checks for overflow conditions.
func reverse2(x int) int {
	str := strconv.Itoa(x)
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '-' {
			reversedStr = "-" + reversedStr
		} else {
			reversedStr += string(str[i])
		}
	}
	result, err := strconv.Atoi(reversedStr)
	if err != nil || result > 1<<31-1 || result < -1<<31 {
		return 0
	}
	return result
}
