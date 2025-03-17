func divideArray(nums []int) bool {
    n := len(nums)
    numToCount := make([]byte, 500)
    countOdd := 0
    for i := 0; i < n; i++ {
        numToCount[nums[i]-1] += 1
        if numToCount[nums[i]-1] % 2 != 0 {
            countOdd++
        } else {
            countOdd--
        }
    }
    return countOdd == 0
}
