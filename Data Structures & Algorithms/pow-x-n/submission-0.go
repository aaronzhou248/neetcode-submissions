func myPow(x float64, n int) float64 {
    if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	negator := 1
	if n < 0 {
		negator = -1
		n *= -1
	}
	temp := float64(1)
	for ; n > 0; n-- {
		temp *= x
	}
	if negator == -1 {
		return 1/temp
	}
	return temp
}
