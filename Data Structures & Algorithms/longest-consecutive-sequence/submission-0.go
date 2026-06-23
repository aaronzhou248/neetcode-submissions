func longestConsecutive(nums []int) int {
	dictionary := map[int]bool{} // this will map a number to whether or not we've used it.
	maxLen := 0

	for _, num := range(nums) {
		dictionary[num] = true
	}

	for _, num := range(nums) {
		if isUnused, succ := dictionary[num]; !succ || !isUnused {
			continue
		}
		currLength := 1
		for left := num - 1; left >= -1000000000; left-- {
			if isUnused, succ := dictionary[left]; !succ || !isUnused {
				break
			}
			currLength += 1
		}
		for right := num + 1; right <= 1000000000; right++ {
			if isUnused, succ := dictionary[right]; !succ || !isUnused {
				break
			}
			currLength += 1
		}

		if currLength > maxLen {
			maxLen = currLength
		}
	}

	return maxLen
}
