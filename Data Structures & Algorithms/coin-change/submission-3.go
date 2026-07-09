func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
	for i := range amount {
		dp[i+1] = math.MaxInt - 1
	}

	for _, coin := range coins { // 1 2 5
		for i := range amount + 1 { // 1 1 1 1 1 1 1 1 1 1 1
			j := i - coin
			if j >= 0 {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt - 1 {
		return -1
	}
	return dp[amount]
}
