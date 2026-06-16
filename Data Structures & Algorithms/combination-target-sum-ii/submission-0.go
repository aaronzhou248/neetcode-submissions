func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	if len(candidates) == 0 {
		return [][]int{}
	}
	if len(candidates) == 1 {
		if candidates[0] == target {
			return [][]int{{candidates[0]}}
		} else {
			return [][]int{}
		}
	}

	result := [][]int{}
	current := []int{}

	var backtrack func(start int, last int)
	backtrack = func(start int, last int) {
		if start > 0 && candidates[start] == candidates[start-1] && last != start-1 {
			return
		}

		sum := 0
		for _, num := range current {
			sum += num
		}
		if sum == target {
			perm := make([]int, len(current))
			copy(perm, current)
			result = append(result, perm)
		}

		for i := start; i < len(candidates)-1; i++ {
			current = append(current, candidates[i+1])
			backtrack(i+1, start)
			current = current[:len(current)-1]
		}
	}
	backtrack(-1, 0)
	return result
}