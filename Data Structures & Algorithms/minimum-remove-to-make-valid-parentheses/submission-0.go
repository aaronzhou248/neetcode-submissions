func minRemoveToMakeValid(s string) string {
	lcount, i := 0, 0
	sa := []rune(s)
	for i < len(s) {
		if sa[i] == '(' {
			lcount++
		} else if sa[i] == ')' {
			if lcount == 0 {
				sa[i] = '!'
			} else {
				lcount--
			}
		}
		i++
	}
	i--
	
	for lcount > 0 {
		if sa[i] == '(' {
			sa[i] = '!'
			lcount--
		}
		i--
	}

	var b strings.Builder
	for _, r := range sa {
		if r == '!' {
			continue
		}
		fmt.Fprintf(&b, "%c", r)
	}
	return b.String()
}
