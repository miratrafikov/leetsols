func subarraySum(nums []int, k int) int {
    subarraysCount := 0
    sumToTimesItOccured := map[int]int{0: 1}
    currentSum := 0
    for i := 0; i < len(nums); i++ {
        currentSum += nums[i]
        if sumOccurences, found := sumToTimesItOccured[currentSum - k]; found {
            subarraysCount += sumOccurences
        }
        sumToTimesItOccured[currentSum]++
    }
    return subarraysCount
}
