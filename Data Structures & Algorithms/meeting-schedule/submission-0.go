/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

type MIH []Interval

func (h MIH) Len() int { return len(h) }
func (h MIH) Less(i, j int) bool { return h[i].start < h[j].start }
func (h MIH) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MIH) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

func (h *MIH) Pop() interface{} {
	n := len(*h)
	old := (*h)[n-1]
	*h = (*h)[:n-1]
	return old
}

func canAttendMeetings(intervals []Interval) bool {
	h := &MIH{}
	heap.Init(h)
	for _, i := range intervals {
		heap.Push(h, i)
	}
	meetingEndTime := 0
	for len(*h) > 0 {
		i := heap.Pop(h).(Interval)
		if i.start < meetingEndTime {
			return false
		}
		meetingEndTime = i.end
	}
	return true
}
