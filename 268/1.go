func missingNumber(nums []int) int {
    n := len(nums)
    expectedSum := 0
    for i := 0; i <= n; i++ {
        expectedSum += i
    }
    actualSum := 0
    for i := 0; i < n; i++ {
        actualSum += nums[i]
    }
    return expectedSum - actualSum
}
