func moveZeroes(nums []int)  {
    i := -1
    for j := 0; j < len(nums); j++ {
        if nums[j] != 0 {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
}
