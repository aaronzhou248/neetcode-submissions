type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder
	for _, str := range strs {
		b.WriteRune(255)
		for _, r := range str {
			if r == 255 {
				b.WriteRune(255)
				b.WriteRune(255)
			} else {
				b.WriteRune(r)
			}
		}
		b.WriteRune(255)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	opened := false
	output := []string{}
	var current strings.Builder
	for i, r := range encoded {
		if opened {
			if r != 255 {
				current.WriteRune(r)
				continue
			}
			if i <= len(encoded)-2 {
				next := encoded[i+1]
				if next != 255 {
					opened = false
					output = append(output, current.String())
					current.Reset()
				} else {
					current.WriteRune(r)
					i++
				}
			}
		} else {
			opened = true // the value of r should be 255.
		}
	}
	return output
}