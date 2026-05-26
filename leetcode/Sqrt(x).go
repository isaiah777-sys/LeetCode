package leetcode

//Note that no built-in function is allowed to solve this problem, so we use the method below to solve it instead.

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 1, x/2
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
