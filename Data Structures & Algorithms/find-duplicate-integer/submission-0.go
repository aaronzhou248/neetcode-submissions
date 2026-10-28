func findDuplicate(nums []int) int {
	found := map[int]struct{}{}
	for _, num := range nums {
		if _, succ := found[num]; succ {
			return num
		} else {
			found[num] = struct{}{}
		}
	}
	return -1
}