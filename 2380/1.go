func secondsToRemoveOccurrences(s string) int {
    if len(s) == 1 {
        return 0
    }
    bytes := []byte(s)
    secondsPassed := 0
    for doesAtLeastOnePairExist(bytes) {
        secondsPassed++
        i := 0
        for i < len(s) - 1 {
            if bytes[i] == '0' && bytes[i+1] == '1' {
                bytes[i], bytes[i+1] = '1', '0'
                i += 2
            } else {
                i++
            }
        }
    }
    return secondsPassed
}

func doesAtLeastOnePairExist(bytes []byte) bool {
    for i := 0; i < len(bytes) - 1; i++ {
        if bytes[i] == '0' && bytes[i+1] == '1' {
            return true
        }
    }
    return false
}

