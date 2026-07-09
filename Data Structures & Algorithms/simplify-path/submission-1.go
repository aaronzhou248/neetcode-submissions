func simplifyPath(path string) string {
	stack := []string{}

	for _, element := range strings.Split(path, "/") {
		switch element {
			case "", ".":
				continue
			case "..":
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			default:
				stack = append(stack, element)
		}
	}

	return "/" + strings.Join(stack, "/")
}
