func fib(n int) int {
    mem := make([]int, max(n+1, 2))
    mem[0] = 0
    mem[1] = 1
    for i := 2; i <= n; i++ {
        mem[i] = mem[i-1] + mem[i-2]
    }
    return mem[n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
