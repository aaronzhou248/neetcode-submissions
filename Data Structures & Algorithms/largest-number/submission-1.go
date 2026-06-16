import "slices"

func largestNumber(nums []int) string {
    strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	slices.SortFunc(strs, func(a, b string) int {
		ab := a + b
		ba := b + a
		if ab < ba {
			return 1
		}
		if ab > ba {
			return -1
		}
		return 0
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
