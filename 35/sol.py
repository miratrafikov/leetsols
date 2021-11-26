class Solution:
    
    def searchInsertUtil(self, l: int, h: int) -> int:
        if l >= h:
            if self.nums[l] < self.target:
                return l + 1
            else:
                return l
        mid = l + (h - l) // 2
        if self.nums[mid] == self.target:
            return mid
        elif self.nums[mid] > self.target:
            return self.searchInsertUtil(l, mid - 1)
        else:
            return self.searchInsertUtil(mid + 1, h)
    
    def searchInsert(self, nums: List[int], target: int) -> int:
        self.nums = nums
        self.target = target
        return self.searchInsertUtil(0, len(nums) - 1)
