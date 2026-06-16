func climbStairs(n int) int {
    // create an array of all the possible "endpoints" and initialize them all to 0
	// except for the 0th index which is the bottom.
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	stairs := make([]int, n + 1)
	stairs[0] = 1
	stairs[1] = 1
	stairs[2] = 2

	for i := 2; i <= n; i++ {
		stairs[i] = stairs[i - 1] + stairs[i - 2]
	}

	return stairs[n]
}
