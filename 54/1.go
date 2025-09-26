func spiralOrder(matrix [][]int) []int {
    n := len(matrix) // vertical length
    m := len(matrix[0]) // horizontal length

    spiral := make([]int, 0, n*m)

    top := 0
    right := m-1
    bottom := n-1
    left := 0

    for top <= bottom && left <= right {
        for i := left; i <= right; i++ {
            spiral = append(spiral, matrix[top][i])
        }
        top++

        for i := top; i <= bottom; i++ {
            spiral = append(spiral, matrix[i][right])
        }
        right--

        if top <= bottom {
            for i := right; i >= left; i-- {
                spiral = append(spiral, matrix[bottom][i])
            }
            bottom--
        }

        if left <= right {
            for i := bottom; i >= top; i-- {
                spiral = append(spiral, matrix[i][left])
            }
            left++
        }
    }

    return spiral
}
