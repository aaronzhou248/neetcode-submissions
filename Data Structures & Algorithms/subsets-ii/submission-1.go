func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	results := [][]int{}
	current := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		perm := make([]int, len(current))
		copy(perm, current)
		results = append(results, perm)

		for i := start; i < len(nums); i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return results
}