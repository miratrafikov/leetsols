class Solution:
    def climbStairs(self, n: int) -> int:
        if n == 1:
            return 1
        if n == 2:
            return 2
        prev = 1
        curr = 2
        for i in range(2, n - 1):
            prev, curr = curr, prev + curr
        return prev + curr
