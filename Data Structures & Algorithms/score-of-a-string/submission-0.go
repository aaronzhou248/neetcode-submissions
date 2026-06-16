func scoreOfString(s string) int {
	sa := []rune(s)
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		diff := sa[i+1] - sa[i]
		if diff > 0 {
			sum += int(diff)
		} else {
			sum -= int(diff)
		}
	}
	return sum
}