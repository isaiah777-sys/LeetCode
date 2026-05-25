package leetcode

func removeElement(nums []int, val int) int {
	wdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[wdx] = nums[i]
			wdx++
		}
	}
	return wdx
}
