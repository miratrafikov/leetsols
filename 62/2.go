func uniquePaths(m int, n int) int {
    // ctw (cell to ways) is a M x N grid where in each cell is a number of ways to reach it
    ctw := make([][]int, m)
    for i := 0; i < m; i++ {
        ctw[i] = make([]int, n)
    }

    ctw[0][0] = 1

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                continue
            }

            top := 0
            if i >= 1 {
                top = ctw[i-1][j]
            }

            left := 0
            if j >= 1 {
                left = ctw[i][j-1]
            }

            ctw[i][j] = top + left
        }
    }

    return ctw[m-1][n-1]
}
