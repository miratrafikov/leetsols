func lengthOfLongestSubstring(s string) int {
    n := len(s)
    longestWindowSize := 0
    windowStart, windowEnd := 0, 0
    windowSet := [256]bool{}
    for windowEnd < n {
        if windowSet[s[windowEnd]] {
            windowSet[s[windowStart]] = false
            windowStart++
        } else {
            windowSet[s[windowEnd]] = true
            windowEnd++
            windowSize := windowEnd - windowStart
            if windowSize > longestWindowSize {
                longestWindowSize = windowSize
            }
        }
    }
    return longestWindowSize
}
