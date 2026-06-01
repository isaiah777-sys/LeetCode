package leetcode

func subsets(nums []int) [][]int {
	result := [][]int{{}} // Start with the empty subset
	for _, num := range nums {
		// For each existing subset, create a new subset that includes the current number
		for _, subset := range result {
			newSubset := append([]int{}, subset...) // Create a copy of the existing subset
			newSubset = append(newSubset, num)
			result = append(result, newSubset) // Add the new subset to the result
		}
	}
	return result
}
