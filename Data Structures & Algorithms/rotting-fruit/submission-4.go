func orangesRotting(grid [][]int) int {
	y := len(grid)
	x := len(grid[0])

	for day := 2; day <= 102; day++ {
		rottenFruitFound := false
		freshFruitFound := false
		newRottingFruit := false
		
		for i := range y {
			for j := range x {
				current := grid[i][j]

				if current == 1 || current == day + 1 {
					freshFruitFound = true
				}
				if current == day {
					rottenFruitFound = true

					if i > 0 && grid[i-1][j] == 1 {
						grid[i-1][j] = day + 1
						newRottingFruit = true
					}
					if i < y - 1 && grid[i+1][j] == 1 {
						grid[i+1][j] = day + 1
						newRottingFruit = true
					}
					if j > 0 && grid[i][j-1] == 1 {
						grid[i][j-1] = day + 1
						newRottingFruit = true
					}
					if j < x - 1 && grid[i][j+1] == 1 {
						grid[i][j+1] = day + 1
						newRottingFruit = true
					}
				}
			}
		}

		if !freshFruitFound {
			return day - 2
		}

		if day == 2 && !rottenFruitFound {
			return -1
		}

		if !newRottingFruit {
			return -1
		}
	}
	return -1
}
