func longestPalindrome(s string) string {
    maxLength, maxString := 0, ""
	sarr := []rune(s)
	n := len(s)

	for i := range n { // ababd --> 1
		l, r := i, i // -1, 3
		for l >= 0 && r < n {
			if sarr[l] == sarr[r] { // a == a
				l--
				r++
			} else {
				break
			}
		}
		if length := r - l - 1; length > maxLength {
			maxLength = length
			maxString = string(sarr[l+1:r])
		}
		l, r = i, i+1 // even length
		for l >= 0 && r < n {
			if sarr[l] == sarr[r] {
				l--
				r++
			} else {
				break
			}
		}
		if length := r - l - 1; length > maxLength {
			maxLength = length
			maxString = string(sarr[l+1:r])
		}
	}

	return maxString
}
