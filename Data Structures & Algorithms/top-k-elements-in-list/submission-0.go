type IIT struct {
	Num  int
	Freq int
}

type MinIITH []IIT

func (h MinIITH) Len() int {
	return len(h)
}

func (h MinIITH) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}

func (h MinIITH) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinIITH) Push(x interface{}) {
	*h = append(*h, x.(IIT))
}

func (h *MinIITH) Pop() interface{} {
	length := h.Len()
	temp := (*h)[length-1]
	(*h) = (*h)[:length-1]
	return temp
}

func topKFrequent(nums []int, k int) []int {
	frequencyGraph := map[int]int{} // this counts the # of occurences for each number
	h := &MinIITH{}
	heap.Init(h)

	// get the frequencies
	for _, num := range nums {
		frequency, found := frequencyGraph[num]
		if !found {
			frequencyGraph[num] = 1
		} else {
			frequencyGraph[num] = frequency + 1
		}
	}

	// now use a max heap.
	for num, frequency := range frequencyGraph {
		if h.Len() < k {
			heap.Push(h, IIT{num, frequency}) // still filling slots
		} else if frequency > (*h)[0].Freq {
			heap.Pop(h)                       // evict the smallest
			heap.Push(h, IIT{num, frequency}) // admit the larger one
		}
		// num >= root? Not top-K material, skip.
	}

	// extract out the value from the heap
	output := []int{}
	for h.Len() > 0 {
		top := heap.Pop(h)
		output = append(output, top.(IIT).Num)
	}

	return output
}