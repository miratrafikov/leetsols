func divideArray(nums []int) bool {
    n := len(nums)
    numToCount := make(map[int]int, n)
    for i := 0; i < n; i++ {
        numToCount[nums[i]] += 1
    }
    for _, num := range numToCount {
        if num % 2 != 0 {
            return false
        }
    }
    return true
}
