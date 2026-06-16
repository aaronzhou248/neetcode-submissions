func hasDuplicate(nums []int) bool {
    appeared := map[int]struct{}{}

	for _, num := range(nums) {
		_, succ := appeared[num]
		if succ {
			return true
		}
		appeared[num] = struct{}{}
	}
	
	return false
}
