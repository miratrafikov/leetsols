class Solution:
    
    def searchInsertUtil(self, nums: List[int], target: int) -> int:
        answer = 0
        l = 0
        h = len(nums) - 1
        while l <= h:
            mid = l + (h - l) // 2
            if nums[mid] == target:
                answer = mid
                break
            elif nums[mid] > target:
                h = mid - 1
                answer = l
            else:
                l = mid + 1
                answer = l
        return answer
    
    def searchInsert(self, nums: List[int], target: int) -> int:
        return self.searchInsertUtil(nums, target)
