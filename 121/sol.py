class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        n = len(prices)

        min, max = [None] * n, [None] * n
        min[0] = prices[0]
        max[n - 1] = prices[n - 1]

        for i in range(n - 2, -1, -1):
            if prices[i] > max[i + 1]:
                max[i] = prices[i]
            else:
                max[i] = max[i + 1]

        for i in range(1, n):
            if prices[i] < min[i - 1]:
                min[i] = prices[i]
            else:
                min[i] = min[i - 1]

        max_profit = 0

        for i in range(n):
            if max[i] - min[i] > max_profit:
                max_profit = max[i] - min[i]

        return max_profit
