func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    var duplicate, missing int
    numberToCount := make([]int, n * n + 1)
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            number := grid[r][c]
            numberToCount[number]++
            if numberToCount[number] == 2 {
                duplicate = number
            }
        }
    }
    for number := 1; number <= n * n; number++ {
        if numberToCount[number] == 0 {
            missing = number
            break
        }
    }
    return []int{duplicate, missing}
}
