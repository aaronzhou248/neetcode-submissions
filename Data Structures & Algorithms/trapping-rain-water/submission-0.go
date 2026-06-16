func trap(height []int) int {
	l := len(height)
	lMax, rMax := make([]int, l), make([]int, l)

	li, ri := 0, 0
	for i := range l {
		if height[i] > li {
			li = height[i]
		}
		lMax[i] = li
		if height[l-i-1] > ri {
			ri = height[l-i-1]
		}
		rMax[l-i-1] = ri
	}

	output := 0
	for i := range l {
		minHeight := min(rMax[i], lMax[i])
		if minHeight > height[i] {
			output += minHeight - height[i]
		}
	}
	return output
}