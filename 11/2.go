func maxArea(height []int) int {
    maxVolume := 0
    left := 0
    right := len(height) - 1
    for left < right {
        columnHeight := min(height[left], height[right])
        columnWidth := right - left
        columnVolume := columnHeight * columnWidth
        maxVolume = max(maxVolume, columnVolume)
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxVolume
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
