func maxSubArray(nums []int) int {
    n := len(nums)
    indexToMaxSubarrayAtIndex := make([]int, n)
    indexToMaxSubarrayAtIndex[0] = nums[0]
    maxSubarray := indexToMaxSubarrayAtIndex[0]
    for i := 1; i < n; i++ {
        indexToMaxSubarrayAtIndex[i] = max(
            nums[i],
            indexToMaxSubarrayAtIndex[i-1] + nums[i],
        )
        maxSubarray = max(
            maxSubarray,
            indexToMaxSubarrayAtIndex[i],
        )
    }
    return maxSubarray
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
