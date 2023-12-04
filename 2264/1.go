func largestGoodInteger(num string) string {
    maxGroupDigit := -1
    bytes := []byte(num)
    for i := 0; i < len(bytes) - 2; i++ {
        groupDigit := int(bytes[i] - '0')
        if groupDigit <= maxGroupDigit {
            continue
        }
        if bytes[i + 1] == bytes[i] && bytes[i + 2] == bytes[i] {
            maxGroupDigit = groupDigit
            i += 2
        }
    }
    if maxGroupDigit == -1 {
        return ""
    }
    digitAsChar := byte('0' + maxGroupDigit)
    return string([]byte{digitAsChar, digitAsChar, digitAsChar})
}
