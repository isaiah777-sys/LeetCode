package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs[1:] {
		for len(prefix) > 0 && str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

// This is not the solution, but it's just a program to print the duplicates in the list.
func printDuplicates(nums []int) []int {
	duplicates := []int{}
	for _, number := range nums {
		if number > len(nums) {
			continue
		}
		if nums[number-1] < 0 {
			duplicates = append(duplicates, number)
		} else {
			nums[number-1] = -nums[number-1]
		}
	}
	return duplicates
}
