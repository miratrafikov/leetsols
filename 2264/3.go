func largestGoodInteger(num string) string {
    maxGroupDigit := ""
    for i := 0; i < len(num) - 2; i++ {
        if num[i] == num[i + 1] && num[i] == num[i + 2] {
            groupDigit := string(num[i])
            if groupDigit > maxGroupDigit {
                maxGroupDigit = groupDigit
            }
        }
    }
    return maxGroupDigit + maxGroupDigit + maxGroupDigit
}
