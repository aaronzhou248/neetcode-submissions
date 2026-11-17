func climbStairs(n int) int {
    dp := make([]int, n+1)
	// setting the base cases.
	dp[0] = 1
	dp[1] = 1
	for i := range(n) - 1 {
		dp[i+2] = dp[i] + dp[i+1]
	}
	return dp[n]
}
