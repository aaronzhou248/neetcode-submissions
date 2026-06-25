func rob(nums []int) int {
	switch len(nums) {
		case 0:
			return 0
		case 1:
			return nums[0]
		case 2:
			return max(nums[0], nums[1])
	}

	n := len(nums)
	type tup struct{
		this int
		last int
	}

	sol := make([]tup, n)

	for i, num := range nums {
		if i == 0 {
			sol[i] = tup{num, 0}
			continue
		} else if i == 1 {
			sol[i] = tup{num, nums[0]}
			continue
		}

		sol[i] = tup{num + max(sol[i-2].this, sol[i-2].last), sol[i-1].this}
	}

	return max(sol[n-1].this, sol[n-1].last)
}
//   2,     3,    7
// 1, 2, 3, 3, 5, 7
// 1, 2, 4, 5, 


// 1, 5, 1, 1, 7
// 1, 5, 6, 6, 

// 1,   5,   1,   1,   7
// 1,0  5,1  2,5  6,5  12,5

// 1,   2,   3,   1
// 1,0  2,1  4,2  3,4