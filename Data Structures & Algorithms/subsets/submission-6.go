func subsets(nums []int) [][]int {
	output := [][]int{}
	var dfs func(current []int, index int)
	dfs = func(current []int, index int) {
		if index == len(nums) {
			copied := make([]int, len(current))
			copy(copied, current)
			output = append(output, copied)
			return
		}

		dfs(append(current, nums[index]), index+1)
		dfs(current, index+1)
	}
	dfs([]int{}, 0)
	return output
}