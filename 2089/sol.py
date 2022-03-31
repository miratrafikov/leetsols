class Solution:
    NOT_FOUND = -1
    
    def find_leftmost_target(self, nums: List[int], target: int) -> int:
        leftmost_target = self.NOT_FOUND
        l, r = 0, len(nums) - 1
        while l <= r:
            mid = l + (r - l) // 2
            if nums[mid] == target:
                leftmost_target = mid
                r = mid - 1
            elif nums[mid] > target:
                r = mid - 1
            else:
                l = mid + 1
        return leftmost_target
        
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums = sorted(nums)
        leftmost_target = self.find_leftmost_target(nums, target)
        if leftmost_target == self.NOT_FOUND:
            return []
        result = [leftmost_target]
        for i in range(leftmost_target + 1, len(nums)):
            if nums[i] != target:
                return result
            result.append(i)
        return result
