func hammingDistance(x int, y int) int {
    mask := 1
    count := 0
    for i := 0; i < 31; i++ {
        if x & mask != y & mask {
            count++
        }
        mask <<= 1
    }
    return count
}
