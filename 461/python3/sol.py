class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        mask = 1
        count = 0
        for i in range(31):
            if x & mask != y & mask:
                count += 1
            mask <<= 1
        return count
