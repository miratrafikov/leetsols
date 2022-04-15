class Solution:
    def binary_search(self, l: int, r: int, target: int) -> int:
        while r >= l:
            mid = l + (r - l) // 2
            if self.nums[mid] == target:
                return mid
            elif self.nums[mid] > target:
                r = mid - 1
            else:
                l = mid + 1
        return -1
    
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        self.nums = nums
        for i in range(len(nums)):
            if nums[i] == target and i < len(nums) - 1 and nums[i + 1] == nums[i]:
                return [i + 1, i + 2]
            pos_pair = self.binary_search(i + 1, len(nums) - 1, target - nums[i])
            if pos_pair != -1:
                return [i + 1, pos_pair + 1]
