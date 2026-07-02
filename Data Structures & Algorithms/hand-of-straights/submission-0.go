import "slices"

func isNStraightHand(hand []int, groupSize int) bool {
    slices.Sort(hand)

	// try to remove the element at i. If it is not the "next"
	// number and not the "default" number i.e. -1, then increment
	// i until hand size is met, at which point, return i to 0 and
	// target to -1
	// 3, 5, 6, 7
	target, i, handSize := -1, 0, 0 // -1 0 0
	for len(hand) > 0 && i < len(hand) {
		current := hand[i] // 4
		if target == -1 || current == target {
			handSize++
			target = current+1
			hand = append(hand[:i], hand[i+1:]...)
		} else {
			i++
		}
		if handSize == groupSize {
			handSize = 0
			target = -1
			i = 0
		}
	}

	return len(hand) == 0 && handSize == 0
}
