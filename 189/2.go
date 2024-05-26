func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

func reverse(nums []int) {
    mid := len(nums) / 2
    for i, j := 0, len(nums)-1; i < mid; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
