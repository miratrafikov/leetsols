class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_to_index = {}
        result = [0, 1]
        if len(nums) == 2:
            return result
        for i, num in enumerate(nums):
            remainder = target - num
            if remainder in num_to_index:
                result = [num_to_index[remainder], i]
                break
            num_to_index[num] = i
        return result

