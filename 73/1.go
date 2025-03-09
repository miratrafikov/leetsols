func setZeroes(matrix [][]int)  {
    n := len(matrix)
    m := len(matrix[0])
    rowsToReset := make(map[int]struct{}, n)
    columnsToReset := make(map[int]struct{}, m)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if matrix[i][j] == 0 {
                if _, found := rowsToReset[i]; !found {
                    rowsToReset[i] = struct{}{}
                }
                if _, found := columnsToReset[j]; !found {
                    columnsToReset[j] = struct{}{}
                }
            }
        }
    }
    for row, _ := range rowsToReset {
        for i := 0; i < m; i++ {
            matrix[row][i] = 0
        }
    }
    for column, _ := range columnsToReset {
        for i := 0; i < n; i++ {
            matrix[i][column] = 0
        }
    }
}
