func findMaxAverage(nums []int, k int) float64 {
    var maxSubarraySum int
    var currentSubarraySum int
    for i := 0; i < k; i++ {
        currentSubarraySum += nums[i]
    }
    maxSubarraySum = currentSubarraySum
    for i := k; i < len(nums); i++ {
        currentSubarraySum += nums[i] - nums[i-k]
        maxSubarraySum = max(maxSubarraySum, currentSubarraySum)
    }
    return float64(maxSubarraySum)/float64(k)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
