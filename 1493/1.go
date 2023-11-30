func longestSubarray(nums []int) int {
    i, j := 0, 0
    maxLength := 0
    doesCurrentSequenceContainOneZero := false
    for j < len(nums) {
        if nums[j] == 1 {
            j++
            windowSize := j - i
            if doesCurrentSequenceContainOneZero {
                windowSize--
            }
            if windowSize > maxLength {
                maxLength = windowSize
            }
        } else {
            if !doesCurrentSequenceContainOneZero {
                doesCurrentSequenceContainOneZero = true
                j++
            } else {
                if nums[i] == 0 {
                    doesCurrentSequenceContainOneZero = false
                }
                i++
            }
        }
    }
    if maxLength == len(nums) {
        maxLength--
    }
    return maxLength
}
