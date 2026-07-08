func minEatingSpeed(piles []int, h int) int { // 1 4 3 2, 9
	sum, maxPileSize := 0, 0
	for _, pile := range piles {
		sum += pile
		if pile > maxPileSize {
			maxPileSize = pile
		}
	}
	average := sum / h // 10 / 9 = 1
	l, r := max(1, average), maxPileSize // 1, 4
	minK := maxPileSize // 4

	for l <= r { // 1 <= 4
		k := l + (r - l) / 2 // 1 + 1 == 2
		time := 0
		for _, pile := range piles {
			time += (pile + k - 1) / k
			// 2, 6, 5, 4 --> 1, 3, 2, 2 == 8
		}

		if time > h { // 8 > 9?
			l = k + 1
			continue
		}
		r = k - 1 // 1
		if k < minK { // k < 4
			minK = k // 2
		}
	}
	return minK
}
