func sortArray(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        original := nums[i]
        var j int
        for j = i - 1; j >= 0 && nums[j] > original; j-- {
            nums[j + 1] = nums[j]
        }

        nums[j + 1] = original
    }

    return nums
}
