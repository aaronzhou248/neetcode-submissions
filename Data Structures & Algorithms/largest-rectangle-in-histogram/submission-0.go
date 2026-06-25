type Bar struct {
	h int
	i int
}

type MinBarHeap []Bar

func (h MinBarHeap) Len() int { return len(h) }
func (h MinBarHeap) Less(i, j int) bool { return h[i].h < h[j].h }
func (h MinBarHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinBarHeap) Push(x any) {
	*h = append(*h, x.(Bar))
}
func (h *MinBarHeap) Pop() any {
	n := len(*h)
	temp := (*h)[n-1]
	*h = (*h)[:n-1]
	return temp
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	bh := &MinBarHeap{}
	heap.Init(bh)

	ranges := []map[int]struct{}{}
	r := map[int]struct{}{}

	for i, h := range heights {
		heap.Push(bh, Bar{h, i})
		r[i] = struct{}{}
	}

	ranges = append(ranges, r)

	for len(*bh) > 0 {
		minBar := heap.Pop(bh).(Bar)
		minHeight := minBar.h
		var ourRange map[int]struct{}
		for i, r := range ranges {
			if _, succ := r[minBar.i]; succ {
				ourRange = r
				if i < len(ranges) - 1 {
					ranges = append(ranges[:i], ranges[i+1:]...)
				} else {
					ranges = ranges[:i]
				}
				break
			}
		}
		// there is no case where ourRange should be nil.
		width := len(ourRange)

		if tempArea := width * minHeight; tempArea > maxArea {
			maxArea = tempArea
		}

		left, right := map[int]struct{}{}, map[int]struct{}{}
		for key, _ := range ourRange {
			if key < minBar.i {
				left[key] = struct{}{}
			}
			if key > minBar.i {
				right[key] = struct{}{}
			}
		}
		ranges = append(ranges, left)
		ranges = append(ranges, right)
	}

	return maxArea
}
