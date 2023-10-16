func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    previousRow := getRow(rowIndex-1)
    currentRow := make([]int, rowIndex+1)
    currentRow[0], currentRow[rowIndex] = 1, 1
    for i := 1; i < rowIndex; i++ {
        currentRow[i] = previousRow[i-1] + previousRow[i]
    }
    return currentRow
}
