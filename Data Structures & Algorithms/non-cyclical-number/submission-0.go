func isHappy(n int) bool {
    nums := map[int]struct{}{}

	for {
		_, inNums := nums[n]
		if n == 1 {
			return true
		}
		if inNums {
			return false
		}
		nums[n] = struct{}{}

		temp := 0
		for n > 0 {
			temp += (n % 10) * (n % 10)
			n /= 10
		}

		n = temp
	}
}
