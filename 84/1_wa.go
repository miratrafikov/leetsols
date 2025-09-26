// Idea: from left to right find minimum height, then compute rectangle area = minheight * distance from l to r.
// Move l pointer or right pointer depdending on which column is smaller.
func largestRectangleArea(heights []int) int {
    n := len(heights)
    answer := 0
    l := 0
    r := n-1
    for l <= r {
        minHeight := heights[l]
        for i := l+1; i <= r; i++ {
            if heights[i] < minHeight {
                minHeight = heights[i]
            }
        }
        maxArea := minHeight * (r - l + 1)
        if maxArea > answer {
            answer = maxArea
        }
        if heights[l] < heights[r] {
            l++
        } else {
            r--
        }
    }
    return answer
}
