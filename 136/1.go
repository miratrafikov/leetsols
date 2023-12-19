func singleNumber(nums []int) int {
    movingXOR := 0
    for _, num := range nums {
        movingXOR ^= num
    }
    return movingXOR
}
