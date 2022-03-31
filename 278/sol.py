# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        l, r = 1, n
        smallest_bad = n
        while r >= l:
            mid = l + (r - l) // 2
            is_bad = isBadVersion(mid)
            if is_bad:
                smallest_bad = mid
                r = mid - 1
            else:
                l = mid + 1
        return smallest_bad
