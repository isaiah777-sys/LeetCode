package leetcode

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		currentArea := (right - left) * min(height[left], height[right])
		if currentArea > maxArea {
			maxArea = currentArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// Another alternative way to solve the problem:
func maxArea2(height []int) int {
	start, end, max := 0, len(height)-1, 0
	for start < end {
		area := (end - start) * min(height[start], height[end])
		if area > max {
			max = area
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return max
}
