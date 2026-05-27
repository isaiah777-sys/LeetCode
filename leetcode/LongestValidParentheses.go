package leetcode

func longestValidParentheses(s string) int {
	maxLength := 0
	stack := []int{-1} // Initialize stack with -1 to handle edge case
	for i, char := range s {
		if char == '(' {
			stack = append(stack, i) // Push index of '(' onto stack
		} else {
			stack = stack[:len(stack)-1] // Pop the last index
			if len(stack) == 0 {
				stack = append(stack, i) // Push current index as base for next valid substring
			} else {
				currentLength := i - stack[len(stack)-1] // Calculate current valid substring length
				if currentLength > maxLength {
					maxLength = currentLength // Update max length if current is greater
				}
			}
		}
	}
	return maxLength
}
