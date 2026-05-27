package leetcode

func maxJumps(arr []int, d int) int {
	n := len(arr)
	dp := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) int {
		if dp[i] != 0 {
			return dp[i]
		}
		maxJump := 1
		for j := 1; j <= d; j++ {
			if i-j >= 0 && arr[i-j] < arr[i] {
				maxJump = max(maxJump, 1+dfs(i-j))
			} else {
				break
			}
		}
		for j := 1; j <= d; j++ {
			if i+j < n && arr[i+j] < arr[i] {
				maxJump = max(maxJump, 1+dfs(i+j))
			} else {
				break
			}
		}
		dp[i] = maxJump
		return maxJump
	}

	maxJumps := 0
	for i := 0; i < n; i++ {
		maxJumps = max(maxJumps, dfs(i))
	}
	return maxJumps
}
