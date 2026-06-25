func minCostClimbingStairs(cost []int) int {
    if len(cost) <= 1 {
		return 0
	}

	sol := make([]int, len(cost) + 1)
	for i := 2; i <= len(cost); i++ {
		sol[i] = min((sol[i - 2] + cost[i -2]), (sol[i - 1] + cost[i - 1]))
	}
	return sol[len(cost)]
}
