func missingNumber(nums []int) int {
    n := len(nums)
    sum := 0
    expectedSum := 0
    for i := 0; i < n; i++ {
        sum += nums[i]
        expectedSum += i + 1
    }
    return expectedSum - sum
}
