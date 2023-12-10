func missingNumber(nums []int) int {
    n := len(nums)
    flags := n
    for i := 0; i < n; i++ {
        flags ^= i ^ nums[i] 
    }
    missingNumber := flags
    return missingNumber
}
