func transpose(matrix [][]int) [][]int {
    // Get number of rows and columns of the input matrix.
    // Output matrix rows and columns count are input matrix columns and rows respectively.
    inRows, inColumns := len(matrix), len(matrix[0])
    outRows, outColumns := inColumns, inRows
    // Create an empty output matrix (at the end of the function the matrix will represent
    // the transposed input matrix).
    transposed := make([][]int, outRows)
    for row := 0; row < outRows; row++ {
        transposed[row] = make([]int, outColumns)
    }
    // Fill in the values in the output matrix
    for row := 0; row < outRows; row++ {
        for column := 0; column < outColumns; column++ {
            transposed[row][column] = matrix[column][row]
        }
    }
    return transposed
}
