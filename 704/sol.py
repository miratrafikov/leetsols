class Solution:
    def search(self, nums: List[int], target: int) -> int:
        self.nums = nums
        self.target = target
        return self.binary_search(0, len(nums) - 1)
    
    def binary_search(self, l: int, r: int) -> int:
        if l > r:
            return -1
        mid = l + (r - l) // 2
        if self.nums[mid] == self.target:
            return mid
        if self.nums[mid] > self.target:
            return self.binary_search(l, mid - 1)
        else:
            return self.binary_search(mid + 1, r)
