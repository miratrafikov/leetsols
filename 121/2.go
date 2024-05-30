func maxProfit(prices []int) int {
    n := len(prices)
    maxPriceSeenStartingFromEnd := prices[n-1]
    maxPriceDifference := 0
    for i := n-2; i >= 0; i-- {
        maxPriceSeenStartingFromEnd = max(maxPriceSeenStartingFromEnd, prices[i])
        maxPriceDifference = max(maxPriceDifference, maxPriceSeenStartingFromEnd-prices[i])
    }
    return maxPriceDifference
}
