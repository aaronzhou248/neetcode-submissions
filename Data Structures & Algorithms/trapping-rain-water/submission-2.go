func trap(height []int) int {
	maxHeight := 0
	currentSum := 0
	total := 0
	furthestTower := -1

	for i, h := range height {
		if h >= maxHeight {
			maxHeight = h
			total += currentSum
			currentSum = 0
			furthestTower = i
			continue
		}

		currentSum += maxHeight - h
	}

	maxHeight = 0
	currentSum = 0
	for i := len(height) - 1; i >= furthestTower; i-- {
		h := height[i]
		if h >= maxHeight {
			maxHeight = h
			total += currentSum
			currentSum = 0
			continue
		}
		currentSum += maxHeight - h
	}

	return total
}
