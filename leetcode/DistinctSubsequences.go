package leetcode

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)

	if m < n {
		return 0
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

//Alternative way of solving it and yet simple and efficient

func numDistinct(s string, t string) int {
	dp := make([]int, len(t)+1)
	dp[0] = 1

	for i := 0; i < len(s); i++ {
		for j := len(t) - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[j+1] += dp[j]
			}
		}
	}

	return dp[len(t)]
}
