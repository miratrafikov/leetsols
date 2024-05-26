func rotate(nums []int, k int)  {
    n := len(nums)
    rotated := make([]int, n)
    for i := 0; i < n; i++ {
        rotated[(i+k)%n] = nums[i]
    }
    for i := 0; i < n; i++ {
        nums[i] = rotated[i]
    }
}
