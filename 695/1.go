func maxAreaOfIsland(grid [][]int) int {
    max := 0
    n := len(grid) - 1
    m := len(grid[0]) - 1
    for i := 0; i <= n; i++ {
        for j := 0; j <= m; j++ {
            if grid[i][j] == 1 {
                islandSize := getIslandSize(grid, i, j, n, m)
                if islandSize > max {
                    max = islandSize
                }
            }
        }
    }
    return max
}

func getIslandSize(grid [][]int, x, y, maxX, maxY int) int {
    grid[x][y] = -1
    size := 1
    if x-1 >= 0 && grid[x-1][y] == 1 {
        size += getIslandSize(grid, x-1, y, maxX, maxY)
    }
    if x+1 <= maxX && grid[x+1][y] == 1 {
        size += getIslandSize(grid, x+1, y, maxX, maxY)
    }
    if y-1 >= 0 && grid[x][y-1] == 1 {
        size += getIslandSize(grid, x, y-1, maxX, maxY)
    }
    if y+1 <= maxY && grid[x][y+1] == 1 {
        size += getIslandSize(grid, x, y+1, maxX, maxY)
    }
    return size
}
