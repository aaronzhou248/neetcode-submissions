func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	queue := make([][2]int, 0)
	fresh := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch grid[i][j] {
			case 2:
				queue = append(queue, [2]int{i, j})
			case 1:
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	for len(queue) > 0 && fresh > 0 {
		next := make([][2]int, 0)
		for _, cell := range queue {
			for _, d := range dirs {
				ni, nj := cell[0]+d[0], cell[1]+d[1]
				if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == 1 {
					grid[ni][nj] = 2
					fresh--
					next = append(next, [2]int{ni, nj})
				}
			}
		}
		queue = next
		minutes++
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}