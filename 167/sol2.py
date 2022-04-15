class Solution:    
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        l = 0
        r = len(nums) - 1
        while r > l:
            s = nums[l] + nums[r]
            if s == target:
                return [l + 1, r + 1]
            elif s > target:
                r -= 1
            else:
                l += 1
