func minOperations(nums []int) int {
    occurencesToMinimalOperationsNeeded := mapOccurencesToMinimalOperationsNeeded(len(nums))
    numberToOccurences := mapNumbersToTheirOccurences(nums)
    totalMinimalNumberOfOperations := 0
    for _, occurences := range numberToOccurences {
        if occurencesToMinimalOperationsNeeded[occurences] == -1 {
            return -1
        }
        totalMinimalNumberOfOperations += occurencesToMinimalOperationsNeeded[occurences]
    }
    return totalMinimalNumberOfOperations
}

func mapOccurencesToMinimalOperationsNeeded(n int) []int {
    occurencesToMinimalOperationsNeeded := make([]int, max(n+1, 4))
    occurencesToMinimalOperationsNeeded[0] = 0
    occurencesToMinimalOperationsNeeded[1] = -1
    occurencesToMinimalOperationsNeeded[2] = 1
    occurencesToMinimalOperationsNeeded[3] = 1
    for i := 4; i <= n; i++ {
        operationsIfSubtractedThree := occurencesToMinimalOperationsNeeded[i-3]
        if operationsIfSubtractedThree != -1 {
            occurencesToMinimalOperationsNeeded[i] = operationsIfSubtractedThree + 1
            continue
        }
        operationsIfSubtractedTwo := occurencesToMinimalOperationsNeeded[i-2]
        if operationsIfSubtractedTwo != -1 {
            occurencesToMinimalOperationsNeeded[i] = operationsIfSubtractedTwo + 1
            continue
        }
        occurencesToMinimalOperationsNeeded[i] = -1
    }
    return occurencesToMinimalOperationsNeeded
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func mapNumbersToTheirOccurences(nums []int) map[int]int {
    numberToOccurences := make(map[int]int)
    for _, num := range nums {
        if occurences, found := numberToOccurences[num]; found {
            numberToOccurences[num] = occurences + 1
        } else {
            numberToOccurences[num] = 1
        }
    }
    return numberToOccurences
}
