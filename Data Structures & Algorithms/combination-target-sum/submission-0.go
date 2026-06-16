func combinationSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	current := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		sum := 0
		for _, val := range current {
			sum += val
		}
		if sum == target {
			perm := make([]int, len(current))
			copy(perm, current)
			result = append(result, perm)
		}
		if sum < target {
			for i := start; i < len(nums); i++ {
				current = append(current, nums[i])
				backtrack(i)
				current = current[:len(current)-1]
			}
		}
	}
	backtrack(0)
	return result
}