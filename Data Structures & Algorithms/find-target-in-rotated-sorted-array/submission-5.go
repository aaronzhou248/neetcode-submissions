func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		fl, fr := nums[l], nums[r]
		i := l + (r - l) / 2
		num := nums[i]
		if num == target {
			return i
		}
		
		// determine which half is sorted.
		if fl <= num { // left half is sorted.
			if num > target && target >= fl {
				r = i - 1
				continue
			}
			l = i + 1
		} else {
			if num < target && target <= fr {
				l = i + 1
				continue
			}
			r = i - 1
		}
	}
	return -1
}
