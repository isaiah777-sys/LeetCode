package leetcode

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 0 {
		return 1<<31 - 1
	}
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	negative := (dividend < 0) != (divisor < 0)
	dividend, divisor = abs(dividend), abs(divisor)
	result := 0
	for dividend >= divisor {
		temp, multiple := divisor, 1
		for dividend >= temp<<1 {
			temp <<= 1
			multiple <<= 1
		}
		dividend -= temp
		result += multiple
	}
	if negative {
		return -result
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
