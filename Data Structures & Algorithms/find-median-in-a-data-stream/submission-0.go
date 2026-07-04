type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	old := (*h)[n-1]
	*h = (*h)[:n-1]
	return old
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	old := (*h)[n-1]
	*h = (*h)[:n-1]
	return old
}

type MedianFinder struct {
	// we will be heavy on the min side.
    Min *MinHeap // min heap goes on the right side
	Max *MaxHeap // max heap goes on the left side
}


func Constructor() MedianFinder {
	temp := MedianFinder{}
    temp.Min = &MinHeap{}
	heap.Init(temp.Min)
	temp.Max = &MaxHeap{}
	heap.Init(temp.Max)
	return temp
}


func (this *MedianFinder) AddNum(num int)  {
    if len(*this.Min) < 1 {
		heap.Push(this.Min, num)
		return
	}
	if len(*this.Max) < 1 {
		right := (*this.Min)[0]
		if num > right {
			heap.Pop(this.Min)
			heap.Push(this.Min, num)
			heap.Push(this.Max, right)
		} else {
			heap.Push(this.Max, num)
		}
		return
	}

	l, r, lc, rc := (*this.Max)[0], (*this.Min)[0], len(*this.Max), len(*this.Min)
	if rc <= lc { // net add one to the right side/min heap
		if num < l {
			heap.Pop(this.Max)
			heap.Push(this.Max, num)
			heap.Push(this.Min, l)
		} else {
			heap.Push(this.Min, num)
		}
	} else { // net add one to the left side/max heap
		if num > r {
			heap.Pop(this.Min)
			heap.Push(this.Max, r)
			heap.Push(this.Min, num)
		} else {
			heap.Push(this.Max, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
    lc, lr := len(*this.Min), len(*this.Max)
	if lc == lr {
		return (float64((*this.Min)[0]) + float64((*this.Max)[0])) / 2
	} else {
		return float64((*this.Min)[0])
	}
}
