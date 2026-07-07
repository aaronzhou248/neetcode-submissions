func rob(nums []int) int {
    bestTakingCurrent, bestSkippingCurrent := 0, 0
	for _, num := range nums {
		bestTakingCurrent, bestSkippingCurrent = max(bestTakingCurrent, num+bestSkippingCurrent), bestTakingCurrent
	}
	return bestTakingCurrent
}
