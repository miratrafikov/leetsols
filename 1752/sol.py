class Solution:
    def check(self, nums: List[int]) -> bool:
        already_rotated = False
        for i in range(1, len(nums)):
            if nums[i] < nums[i - 1]:
                if already_rotated:
                    return False
                else:
                    already_rotated = True
        if already_rotated:
            if nums[0] < nums[len(nums) - 1]:
                return False
        return True
