func permute(nums []int) [][]int {
	result := [][]int{}
	current := []int{}
	used := map[int]bool{}

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			perm := make([]int, len(current))
			copy(perm, current)
			result = append(result, perm)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			current = append(current, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			current = current[:len(current)-1]
		}
	}

	backtrack()
	return result
}