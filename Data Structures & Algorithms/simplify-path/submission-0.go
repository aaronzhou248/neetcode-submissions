func simplifyPath(path string) string {
	elements := []string{}

	var sb strings.Builder
	for _, r := range []rune(path) {
		if r == '/' {
			if sb.Len() > 0 {
				elements = append(elements, sb.String())
				sb.Reset()
			}
			continue
		}
		sb.WriteRune(r)
	}
	if sb.Len() > 0 {
		elements = append(elements, sb.String())
	}

	i := 0
	for i < len(elements) {
		current := elements[i]
		if current == "." {
			// we should remove this element.
			if i == 0 {
				elements = elements[1:]
			} else if i + 1 < len(elements) {
				elements = append(elements[:i], elements[i+1:]...)
			} else {
				elements = elements[:i]
			}
		} else if current == ".." {
			if i == 0 {
				elements = elements[1:]
			} else if i + 1 < len(elements) {
				elements = append(elements[:i-1], elements[i+1:]...)
				i-- // decrement i as we removed 2 elements.
			} else {
				elements = elements[:i-1] // no need to decrement, as we've reached the end.
			}
		} else {
			i++
		}
	}

	return "/" + strings.Join(elements, "/")
}
