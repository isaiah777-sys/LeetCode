package leetcode

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (p[j-2] == s[i-1] || p[j-2] == '.'))
			} else {
				dp[i][j] = dp[i-1][j-1] && (p[j-1] == s[i-1] || p[j-1] == '.')
			}
		}
	}
	return dp[m][n]
}

// Let's try using Javascript to solve the same problem
// function isMatch(s, p) {
//     const m = s.length, n = p.length;
//     const dp = Array.from({ length: m + 1 }, () => Array(n + 1).fill(false));
//	 dp[0][0] = true;
//     for (let j = 1; j <= n; j++) {
//         if (p[j - 1] === '*') {
//             dp[0][j] = dp[0][j - 2];
//         }
//     }
//     for (let i = 1; i <= m; i++) {
//         for (let j = 1; j <= n; j++) {
//             if (p[j - 1] === '*') {
//                 dp[i][j] = dp[i][j - 2] || (dp[i - 1][j] && (p[j - 2] === s[i - 1] || p[j - 2] === '.'));
//             } else {
//                 dp[i][j] = dp[i - 1][j - 1] && (p[j - 1] === s[i - 1] || p[j - 1] === '.');
//             }
//         }
//     }
//     return dp[m][n];
// }
