func tribonacci(n int) int {
    mem := make([]int, 3)
    mem[0] = 0
    mem[1] = 1
    mem[2] = 1
    if n <= 2 {
        return mem[n]
    }
    for i := 0; i <= n - 3; i++ {
        mem[2], mem[1], mem[0] = mem[0] + mem[1] + mem[2], mem[2], mem[1]
    }
    return mem[2]
}
