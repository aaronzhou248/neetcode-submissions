func minWindow(s string, t string) string {
	left := 0 // this is the starting index of our "searching" window.
	s_arr := []rune(s)
	currentLetters := map[rune]int{} // maps runes to their count within our window
	targetLetters := map[rune]int{}
	left_window_index := 0
	right_window_index := 0
	length := 1001 // by definition, our maximum string lengths are 1000

	// create the "target"
	for _, r := range t {
		targetLetters[r] += 1
	}

	for i, r := range s {
		// add the letter to current letters
		currentLetters[r] += 1

		for curr := left; curr <= i; curr++ {
			if curr > left {
				currentLetters[s_arr[curr-1]] -= 1
				left = curr
			}

			valid := true
			for target_r, count := range targetLetters {
				if currentLetters[target_r] < count {
					// this means we are no longer satisfying the condition.
					valid = false
					break
				}
			}

			if valid {
				curr_len := i - curr + 1
				if curr_len < length {
					left_window_index = curr
					right_window_index = i
					length = curr_len
				}
			} else {
				break
			}
		}
	}

	if length == 1001 {
		return ""
	}

	return string(s_arr[left_window_index : right_window_index+1])
}