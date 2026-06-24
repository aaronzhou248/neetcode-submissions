func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	mid := (m + n) / 2
	even := (m + n) % 2 == 0

    // these are the left/right points of the position of the
	// separator, not of the actual values.
    left, right := 0, m
	for left <= right {
		i := left + (right - left) / 2
		j := mid - i

		il, ir, jl, jr := math.MinInt, math.MaxInt, math.MinInt, math.MaxInt

		if i > 0 {
			il = nums1[i-1]
		}
		if i < m {
			ir = nums1[i]
		}
		if j > 0 {
			jl = nums2[j-1]
		}
		if j < n {
			jr = nums2[j]
		}

		if il <= jr && ir >= jl {
			// this is our stopping condition.
			r := math.Min(float64(ir), float64(jr))
			if (even) {
				l := math.Max(float64(il), float64(jl))
				return (float64(l) + float64(r)) / 2.0
			}
			return float64(r)
		}

		if ir < jl {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return 0
}

/*
[1, 2, 3, 4] [5, 6, 7]
mid = (m + n) / 2 = 3
i = (m - 1) / 2 = 2
j = mid - i = 1
il = 3
ir = 4
jl = 5
jr = 6
*/
