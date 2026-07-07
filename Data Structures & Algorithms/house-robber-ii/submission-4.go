func robHelp(nums []int) int {
	bc, bp := 0, 0
	for _, num := range nums {
		bc, bp = max(num + bp, bc), bc
	}
	return bc
}

func rob(nums []int) int {
	n := len(nums)
	switch n {
		case 1:
			return nums[0]
		case 2:
			return max(nums[0], nums[1])
	}
    return max(robHelp(nums[:n-1]), robHelp(nums[1:]))
}
