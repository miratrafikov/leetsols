func singleNumber(nums []int) int {
    for _, num := range nums[1:] {
        nums[0] ^= num
    }
    return nums[0]
}
