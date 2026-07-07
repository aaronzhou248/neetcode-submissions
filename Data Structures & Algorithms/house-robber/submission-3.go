func rob(nums []int) int {
	prev2, prev1 := 0, 0 // best for [..i-2], best for [..i-1]
	for _, num := range nums {
		prev2, prev1 = prev1, max(prev1, prev2+num)
	}
	return prev1
}