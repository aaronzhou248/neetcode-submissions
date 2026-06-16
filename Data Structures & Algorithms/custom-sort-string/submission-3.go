func customSortString(order string, s string) string {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r]++
	}

	var result []rune
	// Append chars in order's sequence
	for _, r := range order {
		for counts[r] > 0 {
			result = append(result, r)
			counts[r]--
		}
	}
	// Append remaining chars not in order
	for _, r := range s {
		if counts[r] > 0 {
			result = append(result, r)
			counts[r]--
		}
	}
	return string(result)
}