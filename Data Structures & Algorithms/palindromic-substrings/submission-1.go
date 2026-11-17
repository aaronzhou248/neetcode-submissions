func countSubstrings(s string) int {
    palindromes := 0

	for i := range(len(s)) {
		palindromes += odd(s, i) + even(s, i)
	}

	return palindromes
}

func odd(s string, i int) int {
	n := len(s)
	for radius := range n {
		if i-radius < 0 || i+radius >= n || s[i-radius] != s[i+radius] {
			return radius
		}
	}
	return 1
}
func even(s string, i int) int {
	n := len(s)
	for radius := range n {
		if i-radius < 0 || i+1+radius >= n || s[i-radius] != s[i+radius+1] {
			return radius
		}
	}
	return 0
}


// we can have even and odd length palidromes. odd is easy.

// func odd(s string, i int) int {} -> get the # of palindromes starting from the i-th position by just expanding outwards.
// func even(s string, i int) int {} -> get the # of palindromes starting from the i-th position by just expanding outwards
// for even though, we need to think about it a bit more.
// i.e. let's think about the string "abba" -> we there is only ONE "even" palindrome, so we EITHER need to succeed from index
// 1 or index 2, but ONLY one and not the other.
// For my example, let's say we want to accept index 1. Then we need to have our "starting" palindrome be the index i (1->b) and i + 1 (2->b).
// This would fail for 2, cause then i(2)->b i+1(3)->a