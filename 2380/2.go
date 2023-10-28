func secondsToRemoveOccurrences(s string) int {
    if len(s) == 1 {
        return 0
    }
    bytes := []byte(s)
    secondsPassed := 0
    doesAtLeastOnePairExist := false
    for i := 0; i < len(bytes) - 1; i++ {
        if bytes[i] == '0' && bytes[i+1] == '1' {
            doesAtLeastOnePairExist = true
            break
        }
    }
    for doesAtLeastOnePairExist {
        secondsPassed++
        doesAtLeastOnePairExist = false
        i := 0
        for i < len(s) - 1 {
            if bytes[i] == '0' && bytes[i+1] == '1' {
                bytes[i], bytes[i+1] = '1', '0'
                if !doesAtLeastOnePairExist {
                    if i > 0 && bytes[i-1] == '0' || i+2 < len(s) && bytes[i+1] == '0' && bytes[i+2] == '1' {
                        doesAtLeastOnePairExist = true
                    }
                }
                i += 2
            } else {
                i++
            }
        }
    }
    return secondsPassed
}

