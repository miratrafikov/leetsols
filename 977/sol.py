class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        neg, pos = -1, len(nums)
        squared = []
        for i in range(len(nums)):
            if nums[i] < 0:
                neg = i
            else:
                pos = i
                break
        while neg >= 0 or pos < len(nums):
            if neg >= 0 and pos < len(nums):
                neg_squared = nums[neg] ** 2
                pos_squared = nums[pos] ** 2
                if neg_squared < pos_squared:
                    squared += [neg_squared]
                    neg -= 1
                else:
                    squared += [pos_squared]
                    pos += 1
            elif neg >= 0:
                squared += [nums[neg] ** 2]
                neg -= 1
            else:
                squared += [nums[pos] ** 2]
                pos += 1
        return squared 
