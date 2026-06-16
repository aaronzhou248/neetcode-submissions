func islandsAndTreasure(grid [][]int) {
	// -1 is untraversable water
	// 0 is treasure.
	// INF is traversable land.
	height := len(grid)
	width := len(grid[0])

	// the idea will be to run a flood fill algorithm from every treasure. For every piece of land, write the distance to the treasure. If there is an existing value, take the smaller value.
	// We can do this in place.

	var bfs func(x, y, dist int)
	bfs = func(x, y, dist int) {
		// our stopping conditions are: edges (<0, >= length), water (-1), or other treasure (0) (if it's not the starting point)
		if x < 0 || x >= width || y < 0 || y >= height || grid[y][x] == -1 || (grid[y][x] == 0 && dist != 0) {
			return
		}

		if dist != 0 {
			current := grid[y][x]
			if current <= dist {
				return
			}
			grid[y][x] = dist
		}

		bfs(x-1, y, dist+1)
		bfs(x+1, y, dist+1)
		bfs(x, y-1, dist+1)
		bfs(x, y+1, dist+1)
	}

	for y := range len(grid) {
		for x := range len(grid[0]) {
			if grid[y][x] == 0 {
				bfs(x, y, 0)
			}
		}
	}
}