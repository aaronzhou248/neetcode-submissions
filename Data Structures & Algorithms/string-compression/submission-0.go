func compress(chars []byte) int {
	var left, right int
	l := len(chars)

	for right < l {
		curr := chars[right]
		count := 0

		// inner loop: consume all consecutive occurrences
		for right < l && chars[right] == curr {
			right++
			count++
		}

		// write the character
		chars[left] = curr
		left++

		// write count only if > 1
		if count > 1 {
			s := strconv.Itoa(count)
			for _, r := range s {
				chars[left] = byte(r)
				left++
			}
		}
	}

	return left
}