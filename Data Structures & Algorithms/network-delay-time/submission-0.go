type NodeDist struct { node, dist int }

type MH []NodeDist

func (h MH) Len() int { return len(h) }
func (h MH) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MH) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MH) Push(x any) { *h = append(*h, x.(NodeDist)) }
func (h *MH) Pop() any {
	l := len(*h)
	old := (*h)[l-1]
	*h = (*h)[:l-1]
	return old
}

func networkDelayTime(times [][]int, n int, k int) int {
	// maps the source to a tuple of the destination and its dist
	dists := map[int][][2]int{}
	for _, time := range times {
		dists[time[0]] = append(dists[time[0]], [2]int{time[1], time[2]})
	}

	// only add elements when we've "started" from them. if they're in here
	// that means we've found the shortest path.
	visited := map[int]int{}

	h := &MH{}
	heap.Init(h)
	heap.Push(h, NodeDist{k, 0})

	for h.Len() > 0 {
		curr := heap.Pop(h).(NodeDist)

		// if we've visited this node already, then we've already found the
		// minimum distance and have already added its elements to the heap.
		if _, succ := visited[curr.node]; succ {
			continue
		}

		visited[curr.node] = curr.dist

		for _, neighbor := range dists[curr.node] {
			// only add the neighbor if it isn't already visited.
			if _, succ := visited[neighbor[0]]; !succ {
				heap.Push(h, NodeDist{neighbor[0], neighbor[1] + curr.dist})
			}
		}
	}

	if len(visited) < n {
		return -1
	}

	maxDist := 0
	for _, dist := range visited {
		maxDist = max(dist, maxDist)
	}
	return maxDist
}