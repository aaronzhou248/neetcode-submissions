func maxAreaOfIsland(grid [][]int) int {
    height, width := len(grid), len(grid[0])

    var bfs func(row, column int) int
    bfs = func(row, column int) int {
        if row < 0 || column < 0 || row >= height || column >= width || grid[row][column] == 0 {
            return 0
        }
        grid[row][column] = 0 // mark visited by zeroing out

        return 1 +
            bfs(row-1, column) +
            bfs(row+1, column) +
            bfs(row, column-1) +
            bfs(row, column+1)
    }

    maxSize := 0
    for i := 0; i < height*width; i++ {
        row, column := i/width, i%width
        if grid[row][column] == 1 {
            maxSize = max(maxSize, bfs(row, column))
        }
    }

    return maxSize
}