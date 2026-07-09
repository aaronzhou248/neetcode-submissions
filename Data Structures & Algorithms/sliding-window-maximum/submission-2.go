type MaxTupleHeap [][2]int // the first element is the value, the second element is the index

func(h MaxTupleHeap) Len() int { return len(h) }
func(h MaxTupleHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func(h MaxTupleHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxTupleHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MaxTupleHeap) Pop() interface{} {
	n := len(*h)
	old := (*h)[n-1]
	*h = (*h)[:n-1]
	return old
}

func maxSlidingWindow(nums []int, k int) []int {
    h := &MaxTupleHeap{}
	heap.Init(h)

	for i := range k {
		heap.Push(h, [2]int{nums[i], i})
	}

	output := make([]int, len(nums) - k + 1)
	for i := k - 1; i < len(nums); i++ {
		current := nums[i]
		heap.Push(h, [2]int{current, i})

		// peak at the top, if it's too old, pop until it is not
		top := (*h)[0] // 0 1 2 3 4 5, k = 3, 5 - 3  == 2 <=
		for top[1] <= i - k {
			heap.Pop(h)
			top = (*h)[0]
		}

		output[i - k + 1] = top[0]
	}

	return output
}
