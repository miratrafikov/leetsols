class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        answer = nums[0]
        for i in range(1, len(nums)):
            largest_at_i = max(nums[i] + nums[i - 1], nums[i])
            nums[i] = largest_at_i
            if largest_at_i > answer:
                answer = largest_at_i
        return answer
