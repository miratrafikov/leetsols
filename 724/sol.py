from functools import reduce

class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        sum_r = reduce(lambda a, b: a + b, nums)
        sum_l = 0
        for i in range(len(nums)):
            if i > 0:
                sum_l += nums[i - 1]
            sum_r -= nums[i]
            if sum_l == sum_r:
                return i
        return - 1
        
