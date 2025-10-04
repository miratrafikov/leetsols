func maxArea(height []int) int {
    n := len(height)
    maxwater := 0

    l := 0
    r := n-1

    for l < r {
        w := r - l
        h := min(height[l], height[r])

        maxwater = max(maxwater, w * h)

        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }

    return maxwater
}
