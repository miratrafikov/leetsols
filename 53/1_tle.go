func maxSubArray(nums []int) int {
    n := len(nums)
    maxSum := nums[0]
    for i := 0; i < n; i++ {
        currentSubarraySum := 0
        for j := i; j < n; j++ {
            currentSubarraySum += nums[j]
            if currentSubarraySum > maxSum {
                maxSum = currentSubarraySum
            }
        }
    }
    return maxSum
}
