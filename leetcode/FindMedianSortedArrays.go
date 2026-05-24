package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1) + len(nums2)
	if m%2 == 1 {
		return float64(findKth(nums1, nums2, m/2+1))
	} else {
		return float64(findKth(nums1, nums2, m/2)+findKth(nums1, nums2, m/2+1)) / 2.0
	}
}

func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) > len(nums2) {
		return findKth(nums2, nums1, k)
	}
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return int(math.Min(float64(nums1[0]), float64(nums2[0])))
	}
	i := int(math.Min(float64(len(nums1)), float64(k/2)))
	j := k - i
	if nums1[i-1] < nums2[j-1] {
		return findKth(nums1[i:], nums2, k-i)
	} else {
		return findKth(nums1, nums2[j:], k-j)
	}
}
