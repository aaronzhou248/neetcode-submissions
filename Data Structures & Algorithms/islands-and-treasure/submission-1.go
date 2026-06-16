func islandsAndTreasure(grid [][]int) {
    height := len(grid)
    width := len(grid[0])
    inf := 2147483647

    // Seed the queue with ALL treasure cells at once
    type point struct{ y, x int }
    q := []point{}

    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            if grid[y][x] == 0 {
                q = append(q, point{y, x})
            }
        }
    }

    dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

    // BFS layer by layer — each cell is visited at most once
    for len(q) > 0 {
        curr := q[0]
        q = q[1:]

        for _, d := range dirs {
            ny, nx := curr.y+d[0], curr.x+d[1]
            if ny < 0 || ny >= height || nx < 0 || nx >= width || grid[ny][nx] != inf {
                continue
            }
            grid[ny][nx] = grid[curr.y][curr.x] + 1
            q = append(q, point{ny, nx})
        }
    }
}