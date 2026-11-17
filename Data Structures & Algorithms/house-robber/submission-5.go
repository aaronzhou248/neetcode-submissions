func rob(nums []int) int {
	// cover the base cases.
	if len(nums) == 1 {
		return nums[0]
	}

    dp := make([][]int, len(nums))
	dp[0] = []int{0, nums[0]}

	for i := range len(nums) - 1 {
		dp[i+1] = []int{max(dp[i][0], dp[i][1]), max(nums[i+1]+dp[i][0], dp[i][1])}
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

// 1  1  3  3
// y  n  y  n
// dynamic -> [<money if did not rob current house>, <money if did rob current house (but not necessarily if it's not profitable)>]
// 0,      1,      2,       3
// [[0,1], [1, 1], [1, 4],  [4, 4]]