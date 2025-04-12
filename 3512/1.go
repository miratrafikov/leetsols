func minOperations(nums []int, k int) int {
    n := len(nums)
    sum := 0
    for i := 0; i < n; i++ {
        sum += nums[i]
    }
    return sum % k
}
