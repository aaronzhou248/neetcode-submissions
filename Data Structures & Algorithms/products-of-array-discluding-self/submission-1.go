func productExceptSelf(nums []int) []int {
	n := len(nums)
	lr := make([]int, n)
	rl := make([]int, n)
	lr[0], rl[n-1] = nums[0], nums[n-1]

	for i := range n - 1 {
		lr[i + 1] = lr[i] * nums[i + 1]
		rl[n-i-2] = rl[n-i-1] * nums[n-i-2]
	}

	output := make([]int, n)
	for i := range n {
		current := 1
		if i > 0 {
			current *= lr[i - 1]
		}
		if i < n - 1 {
			current *= rl[i + 1]
		}
		output[i] = current
	}

	return output
}
