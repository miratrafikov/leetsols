func isToeplitzMatrix(matrix [][]int) bool {
    row, col := len(matrix)-1, 0
    for col < len(matrix[0]) {
        if !isDiagonalUniform(matrix, row, col) {
            return false
        }
        if row > 0 {
            row--
        } else {
            col++
        }
    }
    return true
}

func isDiagonalUniform(matrix [][]int, row, col int) bool {
    pivot := matrix[row][col]
    row++
    col++
    for row < len(matrix) && col < len(matrix[0]) {
        if matrix[row][col] != pivot {
            return false
        }
        row++
        col++
    }
    return true
}
