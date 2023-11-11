func moveZeroes(nums []int)  {
    i := 0
    j := 1
    for j < len(nums) {
        if nums[i] != 0 {
            i++
        } else if nums[j] != 0 {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
        j++
    }
}
