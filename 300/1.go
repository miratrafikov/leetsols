func lengthOfLIS(nums []int) int {
    n := len(nums)
    positionToLIS := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        lis := 1
        for j := i + 1; j < n; j++ {
            if nums[i] < nums[j] {
                candidateLIS := 1 + positionToLIS[j]
                if candidateLIS > lis {
                    lis = candidateLIS
                }
            }
        }
        positionToLIS[i] = lis
    }
    maxLIS := 0
    for i := 0; i < n; i++ {
        if positionToLIS[i] > maxLIS {
            maxLIS = positionToLIS[i]
        }
    }
    return maxLIS
}
