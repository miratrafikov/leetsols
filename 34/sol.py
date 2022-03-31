class Solution:
    def find_edge_target(self, l: int, r: int, toFindLeftEdge: bool) -> int:
        edge = -1
        while l <= r:
            mid = l + (r - l) // 2
            if self.nums[mid] == self.target:
                edge = mid
                if toFindLeftEdge:
                    r = mid - 1
                else:
                    l = mid + 1
            else:
                if edge != -1 and abs(mid - edge) == 1:
                    return edge
                else:
                    if self.nums[mid] > self.target:
                        r = mid - 1
                    else:
                        l = mid + 1
        return edge
    
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        self.nums = nums
        self.target = target
        answer = [-1, -1]
        answer[0] = self.find_edge_target(0, len(nums) - 1, True)
        answer[1] = self.find_edge_target(0, len(nums) - 1, False)
        return answer
