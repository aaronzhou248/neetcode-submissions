func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	n := amount // 12
    dp := make([]int, n+1) // 1 2 3 4 5 2 0 0 0 0 0 0 0
	dp[0] = 1

	for i := 1; i <= n; i++ { // 2
		minToTarget := math.MaxInt // 6
		for _, coin := range coins { // 5
			prevIndex := i - coin // 0
			if prevIndex < 0 {
				continue
			}
			prevCount := dp[prevIndex] // 1
			if prevCount == 0 {
				continue
			}

			minToTarget = min(minToTarget, prevCount + 1) // 6, 2
		}
		if minToTarget != math.MaxInt {
			dp[i] = minToTarget
		}
	}

	return dp[n] - 1
}
