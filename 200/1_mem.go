// MLE
type queue [][]int

func (q *queue) Push(row, col int) {
    *q = append(*q, []int{row, col})
}

func (q *queue) Pop() (int, int) {
    v := (*q)[0]
    *q = (*q)[1:]
    return v[0], v[1]
}

func newQueue() queue {
    return queue(make([][]int, 0))
}

func numIslands(grid [][]byte) int {
    rows, cols := len(grid), len(grid[0])
    var islands int
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if grid[row][col] == '1' {
                islands++
                coverIsland(grid, rows, cols, row, col)
            }
        }
    }
    return islands
}

func coverIsland(grid [][]byte, rows, cols, row, col int) {
    q := newQueue()
    q.Push(row, col)
    for len(q) > 0 {
        row, col := q.Pop()
        grid[row][col] = 'x'
        if row - 1 != -1 && grid[row-1][col] == '1' {
            q.Push(row-1, col)
        }
        if row + 1 != rows && grid[row+1][col] == '1' {
            q.Push(row+1, col)
        }
        if col - 1 != -1 && grid[row][col-1] == '1' {
            q.Push(row, col-1)
        }
        if col + 1 != cols && grid[row][col+1] == '1' {
            q.Push(row,col+1)
        }
    }
}
