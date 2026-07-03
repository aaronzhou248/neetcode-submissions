import "slices"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false // must partition evenly
	}

	count := map[int]int{}
	for _, c := range hand {
		count[c]++
	}

	slices.Sort(hand) // gives us ascending order to probe smallest-first
	for _, c := range hand {
		if count[c] == 0 {
			continue // already consumed
		}
		// c is the smallest unconsumed card — it MUST start a group
		for v := c; v < c+groupSize; v++ {
			if count[v] == 0 {
				return false // can't complete the run
			}
			count[v]--
		}
	}
	return true
}