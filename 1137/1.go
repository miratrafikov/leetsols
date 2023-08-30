func tribonacci(n int) int {
    mem := make([]int, max(n+1, 3))
    mem[0] = 0
    mem[1] = 1
    mem[2] = 1
    for i := 3; i <= n; i++ {
        mem[i] = mem[i-3] + mem[i-2] + mem[i-1]
    }
    return mem[n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
