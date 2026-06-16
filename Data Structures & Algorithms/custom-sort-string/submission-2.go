import ("slices")

func customSortString(order string, s string) string {
	om := map[rune]int{}
	for i, r := range order {
		om[r] = i + 1
	}

	s_slice := []rune(s)

	slices.SortFunc(s_slice, func(a, b rune) int {
		a_index := om[a]
		b_index := om[b]
		return a_index - b_index
	})

	return string(s_slice)
}