func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range m+1 {
		dp[i] = make([]int, n+1)
	}

	dp[1][1] = 1

	for i := range m {
		for j := range n {
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j+1] + dp[i+1][j])
		}
	}

	return dp[m][n]
}
