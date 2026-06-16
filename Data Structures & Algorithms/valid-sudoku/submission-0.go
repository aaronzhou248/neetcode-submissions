func isValidSudoku(board [][]byte) bool {
	for i := range 9 {
		nineR, nineC, nineB := map[byte]byte{}, map[byte]byte{}, map[byte]byte{}
		for j := range 9 {
			// column -> [j][i]
			rowValue := board[j][i]
			nineR[rowValue]++

			// row -> [i][j]
			columnValue := board[i][j]
			nineC[columnValue]++

			// boxValue -> [iDiv * 3 + jMod][iMod * 3 + jDiv]
			iDiv := i / 3
			iMod := i % 3
			jDiv := j / 3
			jMod := j % 3
			boxValue := board[iDiv*3+jMod][iMod*3+jDiv]
			nineB[boxValue]++
		}
		keys := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
		for _, key := range keys {
			if key == '.' {
				continue
			}
			if nineR[key] > 1 {
				return false
			}
			if nineC[key] > 1 {
				return false
			}
			if nineB[key] > 1 {
				return false
			}
		}
	}

	return true
}