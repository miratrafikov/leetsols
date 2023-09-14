func maxArea(height []int) int {
    max := 0
    for i := 0; i < len(height) - 1; i++ {
        for j := i + 1; j < len(height); j++ {
            columnHeight := min(height[i], height[j])
            columnWidth := j - i
            columnVolume := columnHeight * columnWidth
            if columnVolume > max {
                max = columnVolume
            }
        }
    }
    return max
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
