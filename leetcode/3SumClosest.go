package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0 // Not enough numbers to form a triplet
	}
	sort.Ints(nums)                           // Sort the array to use two-pointer technique
	closestSum := nums[0] + nums[1] + nums[2] // Initialize closest sum with the first triplet
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			if abs(currentSum-target) < abs(closestSum-target) {
				closestSum = currentSum // Update closest sum if current is closer to target
			} else if currentSum < target {
				left++ // Move left pointer to increase sum
			} else {
				right-- // Move right pointer to decrease sum
			}
		}
	}
	return closestSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
