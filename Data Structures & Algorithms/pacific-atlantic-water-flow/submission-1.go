func pacificAtlantic(heights [][]int) [][]int {
	dirs := [][]int{ {0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	y := len(heights)
	x := len(heights[0])
    pacific := make([][]bool, y)
	atlantic := make([][]bool, y)
	for i := range y {
		pacific[i] = make([]bool, x)
		atlantic[i] = make([]bool, x)
	}

	pE := [][]int{}
	aE := [][]int{}
	for i := range x {
		pacific[0][i] = true
		atlantic[y - 1][i] = true
		pE = append(pE, []int{0, i})
		aE = append(aE, []int{y - 1, i})
	}
	for i := range y - 1 {
		pacific[i + 1][0] = true
		atlantic[i][x - 1] = true
		pE = append(pE, []int{i + 1, 0})
		aE = append(aE, []int{i, x - 1})
	}

	for len(pE) > 0 {
		coord := pE[0]
		cY, cX := coord[0], coord[1]

		for _, direction := range dirs {
			nY, nX := cY + direction[0], cX + direction[1]

			if nY < 0 || nY >= y || nX < 0 || nX >= x {
				continue
			}

			if heights[nY][nX] >= heights[cY][cX] && !pacific[nY][nX] {
				pacific[nY][nX] = true
				pE = append(pE, []int{nY, nX})
			}
		}

		pE = pE[1:]
	}

	for len(aE) > 0 {
		coord := aE[0]
		cY, cX := coord[0], coord[1]

		for _, direction := range dirs {
			nY, nX := cY + direction[0], cX + direction[1]

			if nY < 0 || nY >= y || nX < 0 || nX >= x {
				continue
			}

			if heights[nY][nX] >= heights[cY][cX] && !atlantic[nY][nX] {
				atlantic[nY][nX] = true
				aE = append(aE, []int{nY, nX})
			}
		}

		aE = aE[1:]
	}

	output := [][]int{}
	for i := range y {
		for j := range x {
			if pacific[i][j] && atlantic[i][j] {
				output = append(output, []int{i,j})
			}
		}
	}

	return output
}
