func plusOne(digits []int) []int {
    l := len(digits)
	carry := 1
	for i := l-1; i>=0; i-- {
		temp := digits[i] + carry
		digits[i] = temp % 10
		carry = temp / 10
	}
	if carry > 0 {
		digits = make([]int, l+1)
		digits[0] = 1
	}
	return digits
}
