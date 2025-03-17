func divideArray(nums []int) bool {
    n := len(nums)
    numToCount := make(map[int]int, n)
    countOdd := 0
    for i := 0; i < n; i++ {
        numToCount[nums[i]] += 1
        if count := numToCount[nums[i]]; count % 2 != 0 {
            countOdd++
        } else {
            countOdd--
        }
    }
    return countOdd == 0
}
