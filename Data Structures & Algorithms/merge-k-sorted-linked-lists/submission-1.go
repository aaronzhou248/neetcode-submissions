/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // Going to use a min heap to store ListNodes. We are
 // going to define it so we handle the case of an empty ListNode.
 // That means if any value is nil, we go the other way.

type MinNodeHeap []*ListNode

func (h MinNodeHeap) Len() int {
	return len(h)
}

func (h MinNodeHeap) Less(x, y int) bool {
	left := h[x]
	right := h[y]
	// if left == nil {
	// 	return false
	// } else if right == nil {
	// 	return true
	// }

	return left.Val < right.Val
}

 func (h MinNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
 }

 func (h *MinNodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
 }

 func (h *MinNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinNodeHeap{}
	heap.Init(h)

	preHead := &ListNode{}
	current := preHead

	for _, list := range lists {
		if list == nil {
			continue
		}
		heap.Push(h, list)
	}

	for h.Len() > 0 {
		bottom := heap.Pop(h).(*ListNode)
		current.Next = bottom
		current = bottom

		if bottom.Next != nil {
			heap.Push(h, bottom.Next)
		}
	}

	return preHead.Next
}
