func maxSubArray(nums []int) int {
	best := nums[0]
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		curr = max(nums[i], curr+nums[i])
		best = max(best, curr)
	}
	return best
}