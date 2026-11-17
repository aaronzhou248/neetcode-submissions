func hammingWeight(n int) int {
	ones := 0

	for n > 0 {
		if n % 2 == 1 {
			ones++
		}
		n = n / 2
	}
	
	return ones
}
