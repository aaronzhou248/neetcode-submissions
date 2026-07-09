func longestConsecutive(nums []int) int {
	numMap := map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}

	maxLength := 0
	for _, num := range nums {
		if numMap[num-1] {
			continue
		}

		length := 0
		for {
			if !numMap[num + length] {
				break
			}
			length++
		}

		maxLength = max(maxLength, length)
	}
	return maxLength
}
