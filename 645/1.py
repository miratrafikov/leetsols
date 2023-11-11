class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        n = len(nums)
        encountered = [0]*n
        for i in range(n):
            encountered[nums[i]-1] += 1
        duplicate = None
        missing = None
        i = 0
        while duplicate is None or missing is None:
            if encountered[i] == 0:
                missing = i + 1
            elif encountered[i] == 2:
                duplicate = i + 1
            i += 1
        return [duplicate, missing]

