class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [None] * n * 2
        
        for i in range(n):
            ans[i] = ans[i + n] = nums[i]
        
        return ans
