package leetcode

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	maxNum := make([]int, k)
	for i := 0; i <= k; i++ {
		if i <= len(nums1) && k-i <= len(nums2) {
			merged := merge(maxSubsequence(nums1, i), maxSubsequence(nums2, k-i))
			if compare(merged, 0, maxNum, 0) > 0 {
				copy(maxNum, merged)
			}
		}
	}
	return maxNum
}

func maxSubsequence(nums []int, k int) []int {
	drop := len(nums) - k
	stack := []int{}
	for _, num := range nums {
		for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < num {
			stack = stack[:len(stack)-1]
			drop--
		}
		stack = append(stack, num)
	}
	return stack[:k]
}

func merge(nums1 []int, nums2 []int) []int {
	merged := make([]int, len(nums1)+len(nums2))
	i, j := 0, 0
	for k := 0; k < len(merged); k++ {
		if compare(nums1, i, nums2, j) > 0 {
			merged[k] = nums1[i]
			i++
		} else {
			merged[k] = nums2[j]
			j++
		}
	}
	return merged
}

func compare(nums1 []int, i int, nums2 []int, j int) int {
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] != nums2[j] {
			return nums1[i] - nums2[j]
		}
		i++
		j++
	}
	return (len(nums1) - i) - (len(nums2) - j)
}
