func closestPrimes(left int, right int) []int {
    if right <= 2 {
        return []int{-1,-1}
    }
    areComposites := make([]bool, right+1)
    areComposites[1] = true
    areComposites[2] = false
    for i := 2; i <= right; i++ {
        sequenceTail := i + i
        for sequenceTail <= right {
            areComposites[sequenceTail] = true
            sequenceTail += i
        }
    }
    prevPrime := -1
    currPrime := -1
    minDiff := 60000
    answer := []int{-1, -1}
    for i := left; i <= right; i++ {
        if areComposites[i] == false {
            if currPrime != -1 {
                prevPrime = currPrime
                currPrime = i
                if currPrime - prevPrime < minDiff {
                    answer[0] = prevPrime
                    answer[1] = currPrime
                    minDiff = currPrime - prevPrime
                    if minDiff == 2 {
                        return answer
                    }
                }
            } else {
                if prevPrime == -1 {
                    prevPrime = i
                } else {
                    currPrime = i
                    answer[0] = prevPrime
                    answer[1] = currPrime
                    minDiff = currPrime - prevPrime
                    if minDiff == 2 {
                        return answer
                    }
                }
            }
        }
    }
    return answer
}
