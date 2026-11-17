func countBits(n int) []int {
	output := make([]int, n+1)
	for i := range n + 1 {
		current := i
		ones := 0
		for current > 0 {
			if current % 2 == 1 {
				ones++
			}
			current = current / 2
		}
		output[i] = ones 
	}
	return output
}
