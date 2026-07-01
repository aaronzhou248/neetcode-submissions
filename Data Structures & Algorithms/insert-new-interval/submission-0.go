func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	left, right := newInterval[0], newInterval[1]

	// iterate through all of the existing intervals. They can either:
	// 1. be COMPLETELY before the new interval
	// 2. WITHIN the new interval
	// 3. Completely beyond the new interval
	i, n := 0, len(intervals)

	for i < n && left > intervals[i][1] {
		result = append(result, intervals[i])
		i++
	}
	for i < n && right >= intervals[i][0] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}