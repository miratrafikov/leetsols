func longestSubarray(nums []int) int {
    l, r := 0, 0
    maxLength := 0
    zerosInWindow := 0
    for r < len(nums) {
        if nums[r] == 0 {
            zerosInWindow++
        }
        for zerosInWindow > 1 {
            if nums[l] == 0 {
                zerosInWindow--
            }
            l++
        }
        if r - l > maxLength {
            maxLength = r - l
        }
        r++
    }
    return maxLength
}
