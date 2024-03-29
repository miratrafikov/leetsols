func isToeplitzMatrix(matrix [][]int) bool {
    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i-1][j-1] != matrix[i][j] {
                return false
            }
        }
    }
    return true
}
