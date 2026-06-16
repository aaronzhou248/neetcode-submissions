import (
	"slices"
)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	type Event struct {
		username	string
		timestamp	int
		website		string
	}
	l := len(username)
	events := make([]Event, l)
	for i := range l {
		events[i] = Event{username[i], timestamp[i], website[i]}
	}
	slices.SortFunc(events, func(a, b Event) int {
		if a.username == b.username {
			// this is the same user. Order by timestamp.
			if a.timestamp < b.timestamp {
				return -1
			}
			return 1
		}
		if a.username < b.username {
			return 1
		}
		return -1
	})

	patterns := map[[3]string]int{}
	// we want to iterate everything, less 2
	// as we're looking for triplets
	for i := range l - 2 {
		e0, e1, e2 := events[i], events[i+1], events[i+2]
		if e0.username != e1.username || e1.username != e2.username {
			continue
		}
		pattern := [3]string{e0.website, e1.website, e2.website}
		patterns[pattern]++
	}

	type PatternCount struct {
		pattern [3]string
		count 	int
	}

	pcs := []PatternCount{}
	for key, value := range patterns {
		pcs = append(pcs, PatternCount{key, value})
	}
	slices.SortFunc(pcs, func(a, b PatternCount) int {
		if a.count == b.count {
			for i := 0; i < 3; i++ {
				if a.pattern[i] < b.pattern[i] {
					return -1
				}
				if a.pattern[i] > b.pattern[i] {
					return 1
				}
			}
			return 0
		}
		if a.count < b.count {
			return 1
		}
		return -1
	})

	return pcs[0].pattern[:]
}
