class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        self.sortIntervalsByRightBoundary(intervals)
        maxNumberOfNonOverlappingIntervals = 1
        rightBoundaryOfMostRecentInterval = intervals[0][1]
        for interval in intervals[1:]:
            if interval[0] >= rightBoundaryOfMostRecentInterval:
                rightBoundaryOfMostRecentInterval = interval[1]
                maxNumberOfNonOverlappingIntervals += 1
        return len(intervals) - maxNumberOfNonOverlappingIntervals
    
    def sortIntervalsByRightBoundary(self, intervals: List[List[int]]):
        self.quicksort(intervals, 0, len(intervals)-1)
    
    def quicksort(self, intervals: List[List[int]], l: int, r: int):
        if l == r:
            return
        m = self.partition(intervals, l, r)
        self.quicksort(intervals, l, m)
        self.quicksort(intervals, m+1, r)
    
    def partition(self, intervals: List[List[int]], l: int, r: int) -> int:
        pivot = intervals[l + (r - l) // 2]
        i = l - 1
        j = r + 1
        while True:
            while True:
                i += 1
                if intervals[i][1] >= pivot[1]:
                    break
            while True:
                j -= 1
                if intervals[j][1] <= pivot[1]:
                    break
            if i >= j:
                break
            intervals[i], intervals[j] = intervals[j], intervals[i]
        return j
