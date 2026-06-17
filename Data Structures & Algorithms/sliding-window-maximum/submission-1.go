type IntTuple struct {
	Value int
	Index int
}

type MaxIntTupleHeap []IntTuple

func (h MaxIntTupleHeap) Len() int           { return len(h) }
func (h MaxIntTupleHeap) Less(i, j int) bool { return h[i].Value > h[j].Value }
func (h MaxIntTupleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntTupleHeap) Push(x interface{}) {
	*h = append(*h, x.(IntTuple))
}

func (h *MaxIntTupleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	mitp := &MaxIntTupleHeap{}
	heap.Init(mitp)
	output := []int{}

	for i := range k - 1 {
		heap.Push(mitp, IntTuple{Value: nums[i], Index: i})
	}

	for i := k - 1; i < len(nums); i++ {
		heap.Push(mitp, IntTuple{Value: nums[i], Index: i})
		max := heap.Pop(mitp).(IntTuple)
		// if max is outside the window, we keep popping until it is.
		for max.Index <= i-k {
			max = heap.Pop(mitp).(IntTuple)
		}

		heap.Push(mitp, max)
		output = append(output, max.Value)
	}

	return output
}