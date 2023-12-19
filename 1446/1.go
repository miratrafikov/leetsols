func maxPower(s string) int {
    n := len(s)
    max := 0
    currentChar, currentCount := s[0], 0
    for i := 0; i < n; i++ {
        if s[i] == currentChar {
            currentCount++
            if currentCount > max {
                max = currentCount
            }
        } else {
            currentChar, currentCount = s[i], 1
        }
    }
    return max
}
