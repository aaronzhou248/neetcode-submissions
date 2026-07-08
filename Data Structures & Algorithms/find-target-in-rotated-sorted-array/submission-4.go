func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] { // LEFT half is sorted
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1 // target is in the sorted left half
			} else {
				l = mid + 1 // must be in the right half
			}
		} else { // RIGHT half is sorted
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1 // target is in the sorted right half
			} else {
				r = mid - 1 // must be in the left half
			}
		}
	}
	return -1
}