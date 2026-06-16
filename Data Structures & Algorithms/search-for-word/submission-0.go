func exist(board [][]byte, word string) bool {
	found := false
	// we will store all numbers as x << 16 + y -> 1, 2 -> 0x10000000000000010
	current := map[int]struct{}{} // this will list all coordinates we've visited.

	var backtrack func(indexToSearch, x, y int)
	backtrack = func(indexToSearch, x, y int) {
		if found {
			return
		}
		_, alreadyExists := current[x<<16+y]
		if y > len(board)-1 || y < 0 || x > len(board[y])-1 || x < 0 || board[y][x] != word[indexToSearch] || alreadyExists {
			return
		}

		indexToSearch++
		if indexToSearch == len(word) {
			found = true
			return
		}

		current[x<<16+y] = struct{}{}
		backtrack(indexToSearch, x+1, y)
		backtrack(indexToSearch, x-1, y)
		backtrack(indexToSearch, x, y+1)
		backtrack(indexToSearch, x, y-1)
		delete(current, x<<16+y)
	}
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			backtrack(0, x, y)
		}
	}
	return found
}