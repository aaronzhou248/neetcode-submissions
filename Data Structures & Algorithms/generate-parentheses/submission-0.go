func generateParenthesis(n int) []string {
	result := []string{}
	current := ""

	var backtrack func(open, completed int)
	backtrack = func(open, completed int) {
		if completed == n {
			result = append(result, current)
			return
		}

		if open > 0 {
			current += ")"
			backtrack(open-1, completed+1)
			current = current[:len(current)-1]
		}
		if open+completed < n {
			current += "("
			backtrack(open+1, completed)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, 0)
	return result
}