import "slices"

func subsets(nums []int) [][]int {
	output := [][]int{} // this contains all of the outputs.

	var dfs func(current []int, index int)
	dfs = func(current []int, index int) {
		if index == len(nums) {
			copied := slices.Clone(current)
			output = append(output, copied)
			return
		}
		dfs(current, index+1)
		dfs(append(current, nums[index]), index+1)
	}
	
	dfs([]int{}, 0)
	return output
}