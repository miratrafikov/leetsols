func sortArray(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        for j := i; j > 0 && nums[j - 1] > nums[j]; j-- {
            temp := nums[j]
            nums[j] = nums[j - 1]
            nums[j - 1] = temp
        }
    }

    return nums
}
