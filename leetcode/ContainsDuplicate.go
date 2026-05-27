package leetcode

func containsDuplicate(nums []int) bool {
	appear := make(map[int]bool)
	for _, num := range nums {
		if appear[num] {
			return true
		}
		appear[num] = true
	}
	return false
}
