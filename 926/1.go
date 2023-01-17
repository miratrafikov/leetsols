func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

func minFlipsMonoIncr(s string) int {
    timesOneSeen, flipsNeededToSatisfy := 0, 0
    for _, char := range s {
        if char == '0' {
            if timesOneSeen > 0 {
                flipsNeededToSatisfy += 1
            }
        } else {
            timesOneSeen += 1
        }

        flipsNeededToSatisfy = min(timesOneSeen, flipsNeededToSatisfy)
    }

    return flipsNeededToSatisfy
}
