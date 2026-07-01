func jump(nums []int) int { // [1, 2]
	n := len(nums) // 2
    jumps := make([]int, n) // 0, infty
	for i := 1; i < n; i++ {
		jumps[i] = math.MaxInt
	}

	for i, num := range nums { // 0, 1
		for j := range num { // 0
			if i+j+1 >= n { // 0 + 0 + 1 >= 2?
				break
			}
			jumps[i+j+1] = min(jumps[i+j+1], jumps[i]+1) // jumps[1] = min(infty, 0 + 1) = 2
		}
	}

	return jumps[n-1]
}
