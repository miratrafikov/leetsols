func fib(n int) int {
    if n <= 1 {
        return n
    }
    twoBack := 0
    oneBack := 1
    for i := 2; i < n; i++ {
        twoBack, oneBack = oneBack, twoBack + oneBack
    }
    return twoBack + oneBack
}
