func majorityElement(nums []int) int {
    numberToOccurences := mapNumbersToOccurences(nums)
    var numberWithMostOccurences int = nums[0]
    for k, v := range numberToOccurences {
        if k != numberWithMostOccurences && v > numberToOccurences[numberWithMostOccurences] {
            numberWithMostOccurences = k
        }
    }
    return numberWithMostOccurences
}

func mapNumbersToOccurences(nums []int) map[int]int {
    numberToOccurences := make(map[int]int)
    for _, num := range nums {
        if count, found := numberToOccurences[num]; found {
            numberToOccurences[num] = count + 1
        } else {
            numberToOccurences[num] = 1
        }
    }
    return numberToOccurences
}
