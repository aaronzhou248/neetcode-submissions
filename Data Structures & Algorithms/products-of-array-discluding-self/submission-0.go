func productExceptSelf(nums []int) []int {
	l := len(nums)
	left, right := make([]int, l), make([]int, l)
	lp, rp := 1, 1
	for i := 0; i < l; i++ {
		lp *= nums[i]
		rp *= nums[l-i-1]
		left[i] = lp
		right[l-i-1] = rp
	}
	output := make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			output[i] = right[i+1]
		} else if i == l-1 {
			output[i] = left[i-1]
		} else {
			output[i] = left[i-1] * right[i+1]
		}
	}

	return output
}