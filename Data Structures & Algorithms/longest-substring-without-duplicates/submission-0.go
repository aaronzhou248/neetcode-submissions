func lengthOfLongestSubstring(s string) int {
    lastSeen := map[rune]int{}
    left := 0
    maxLength := 0
 
    for i, r := range []rune(s) {
        if idx, ok := lastSeen[r]; ok && idx >= left {
            left = idx + 1
        }
        lastSeen[r] = i
        if curr := i - left + 1; curr > maxLength {
            maxLength = curr
        }
    }
    return maxLength
 }