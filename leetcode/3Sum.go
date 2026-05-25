package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}
	//We can sort the array to make it easier to find triplets and avoid duplicates
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		// Ignore duplicate elements to avoid duplicate triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// Ignore duplicate elements to avoid duplicate triplets
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
