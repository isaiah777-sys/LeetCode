package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func(start int, path []int, target int)

	backtrack = func(start int, path []int, target int) {
		if target == 0 {
			// Found a valid combination
			combination := append([]int{}, path...)
			result = append(result, combination)
			return
		}
		if target < 0 {
			// Exceeded the target, no valid combination
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, path, target-candidates[i])
			path = path[:len(path)-1] // Backtrack
		}
	}
	backtrack(0, []int{}, target)
	return result
}
