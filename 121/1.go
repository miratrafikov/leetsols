func maxProfit(prices []int) int {
    n := len(prices)
    min := make([]int, n)
    max := make([]int, n)
    min[0] = prices[0]
    max[n - 1] = prices[n - 1]
    for i, j := 1, n - 2; i < n; i, j = i + 1, j - 1 {
        if prices[i] < min[i - 1] {
            min[i] = prices[i]
        } else {
            min[i] = min[i - 1]
        }
        if prices[j] > max[j + 1] {
            max[j] = prices[j]
        } else {
            max[j] = max[j + 1]
        }
    }
    maxDifference := max[n - 1] - min[0]
    for i := 1; i < n; i++ {
        difference := max[i] - min[i]
        if difference > maxDifference {
            maxDifference = difference
        }
    }
    return maxDifference
}
