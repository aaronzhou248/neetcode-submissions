func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i, row := range grid {
		visited[i] = make([]bool, len(row))
	}

	height, width := len(grid), len(grid[0])

	var bfs func(row, column int)
	bfs = func(row, column int) {
		if row < 0 || column < 0 || row >= height || column >= width || visited[row][column] {
			return
		}
		visited[row][column] = true // just mark that we've been here.
		if grid[row][column] == '0' {
			return
		}
		bfs(row-1, column)
		bfs(row+1, column)
		bfs(row, column-1)
		bfs(row, column+1)
	}

	islandCount := 0
	for i := 0; i < height*width; i++ {
		row, column := i/width, i%width
		if !visited[row][column] && grid[row][column] != '0' {
			islandCount++
			bfs(row, column)
		}
	}
	return islandCount
}