func numSpecial(mat [][]int) int {
    rows, columns := len(mat), len(mat[0])
    rowToDegree, columnToDegree := make([]int, rows), make([]int, columns)
    positionsOfOnes := make([][]int, 0)
    for row := 0; row < rows; row++ {
        for column := 0; column < columns; column++ {
            if mat[row][column] == 1 {
                positionsOfOnes = append(positionsOfOnes, []int{row, column})
                rowToDegree[row]++
                columnToDegree[column]++
            }
        }
    }
    specialPositionsCount := 0
    for _, positionOfOne := range positionsOfOnes {
        row, column := positionOfOne[0], positionOfOne[1]
        if rowToDegree[row] == 1 && columnToDegree[column] == 1 {
            specialPositionsCount++
        }
    }
    return specialPositionsCount
}
