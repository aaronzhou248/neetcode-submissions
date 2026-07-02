func maxCoins(nums []int) int {
	n := len(nums)
	b := make([]int, n+2)
	b[0], b[n+1] = 1, 1 // end walls of 1.
	for i, val := range nums {
		b[i+1] = val
	}

	dp := make([][]int, n+2)
	for i, _ := range dp {
		dp[i] = make([]int, n+2)
	}

	// iterate over length, then left (and right), then the "k" balloon to pop
	// minimum length is +2 -> 0 2 has 1 in the middle
	for l:=2; l<n+2; l++ {
		for i := 0; i+l<n+2; i++ {
			j := i+l
			for k := i+1; k<j; 	k++ {
				val := dp[i][k]+dp[k][j]+b[k]*b[i]*b[j]
				if val > dp[i][j] {
					dp[i][j] = val
				}
			}
		}
	}

	return dp[0][n+1]
}

// Greedy does not work.
// func maxCoins(nums []int) int {
// 	n := len(nums)
// 	// pad with 1s on both ends
// 	balloons := make([]int, n+2)
// 	balloons[0], balloons[n+1] = 1, 1
// 	for i := 0; i < n; i++ {
// 		balloons[i+1] = nums[i]
// 	}

// 	// dp[i][j] = max coins from bursting all balloons strictly between i and j
// 	dp := make([][]int, n+2)
// 	for i := range dp {
// 		dp[i] = make([]int, n+2)
// 	}

// 	// length = gap between boundaries; need at least one balloon between
// 	for length := 2; length <= n+1; length++ {
// 		for i := 0; i+length <= n+1; i++ {
// 			j := i + length
// 			for k := i + 1; k < j; k++ { // k = last balloon burst in (i, j)
// 				coins := balloons[i]*balloons[k]*balloons[j] + dp[i][k] + dp[k][j]
// 				if coins > dp[i][j] {
// 					dp[i][j] = coins
// 				}
// 			}
// 		}
// 	}
// 	return dp[0][n+1]
// }