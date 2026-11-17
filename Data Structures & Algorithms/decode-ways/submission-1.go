func numDecodings(s string) int {
	next, next2 := 1, 0 // dp[i+1], dp[i+2]
	for i := len(s) - 1; i >= 0; i-- {
		cur := 0
		if s[i] != '0' {
			cur = next
			if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] <= '6')) {
				cur += next2
			}
		}
		next, next2 = cur, next
	}
	return next
}