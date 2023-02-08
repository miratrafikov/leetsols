func containsNearbyDuplicate(nums []int, k int) bool {
    n := len(nums)
    for i := 0; i < n - 1; i++ {
        for j := 1; i + j < n && j <= k; j++ {
            if nums[i] == nums[i + j] {
                return true
            }
        }
    }

    return false
}
