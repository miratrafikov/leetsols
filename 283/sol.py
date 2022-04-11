class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        pz = 0
        pnz = 1
        while pnz < len(nums):
            if nums[pz] != 0:
                pz += 1
                pnz += 1
            else:
                if nums[pnz] == 0:
                    pnz += 1
                else:
                    nums[pz] = nums[pnz]
                    nums[pnz] = 0
                    pz += 1
                    pnz += 1
