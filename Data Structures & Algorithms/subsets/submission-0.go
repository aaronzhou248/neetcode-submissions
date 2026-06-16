func subsets(nums []int) [][]int {
	result := map[string][]int{} // stores the hashes of the subsets.
	current := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if _, succ := result[fmt.Sprint(current)]; !succ {
			copied := make([]int, len(current))
			copy(copied, current)
			result[fmt.Sprint(current)] = copied
		}

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])

			backtrack(i + 1)

			current = current[:len(current)-1]
		}
	}
	backtrack(0)

	output := [][]int{}
	for _, value := range result {
		output = append(output, value)
	}
	return output
}