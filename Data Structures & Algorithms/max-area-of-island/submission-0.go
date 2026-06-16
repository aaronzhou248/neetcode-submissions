func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]byte, len(grid))
	for i, row := range grid {
		visited[i] = make([]byte, len(row))
	}

	height, width := len(grid), len(grid[0])

	var bfs func(row, column, islandIndex int)
	bfs = func(row, column, islandIndex int) {
		if row < 0 || column < 0 || row >= height || column >= width || visited[row][column] > 0 {
			return
		}
		visited[row][column] = 1 // just mark that we've been here.
		if grid[row][column] == 0 {
			return
		}

		// if we're actually on the piece of land, we're gonna left shift the number.
		visited[row][column] = byte(islandIndex << 1)

		bfs(row-1, column, islandIndex)
		bfs(row+1, column, islandIndex)
		bfs(row, column-1, islandIndex)
		bfs(row, column+1, islandIndex)
	}

	islandCount := 0
	for i := 0; i < height*width; i++ {
		row, column := i/width, i%width
		if visited[row][column] == 0 && grid[row][column] != 0 {
			islandCount++
			bfs(row, column, islandCount)
		}
	}

	islandSizes := make([]int, islandCount)
	for i := 0; i < height*width; i++ {
		row, column := i/width, i%width
		islandIndex := visited[row][column]>>1 - 1
		if islandIndex == 255 {
			continue
		}
		islandSizes[islandIndex]++
	}

	maxSize := 0
	for _, size := range islandSizes {
		if size > maxSize {
			maxSize = size
		}
	}

	return maxSize
}