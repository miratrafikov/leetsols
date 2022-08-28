class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        low = prices[0]
        
        max_profit = 0
        
        for i in range(len(prices)):
            if prices[i] < low:
                low = prices[i]
            else:
                current_profit = prices[i] - low

                if current_profit > max_profit:
                    max_profit = current_profit
        
        return max_profit
