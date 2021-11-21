func rob(nums []int) int {
    if len(nums) > 1 {
        nums[1] = max(nums[0], nums[1])
    }
    if len(nums) > 2 {
        for i := 2; i < len(nums); i++ {
            nums[i] = max(nums[i - 1], nums[i - 2] + nums[i])
        }
    }
    return nums[len(nums) - 1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
