class Solution:
    def findLongestChain(self, pairs: List[List[int]]) -> int:
        self.mergesort(pairs, 0, len(pairs) - 1)
        length = 1
        tail = 0
        for i in range(1, len(pairs)):
            if pairs[i][0] > pairs[tail][1]:
                length += 1
                tail = i
        return length
    
    def mergesort(self, array: List[List[int]], l: int, r: int) -> None:
        if l >= r:
            return
        m = l + (r - l) // 2
        self.mergesort(array, l, m)
        self.mergesort(array, m + 1, r)
        self.merge(array, l, m, r)
    
    def merge(self, array: List[List[int]], l: int, m: int, r: int) -> None:
        left = [0] * (m - l + 1)
        right = [0] * (r - m)
        for i in range(len(left)):
            left[i] = array[l + i]
        for i in range(len(right)):
            right[i] = array[m + 1 + i]
        i = 0
        j = 0
        k = l
        while i < len(left) and j < len(right):
            if left[i][1] < right[j][1]:
                array[k] = left[i]
                i += 1
            else:
                array[k] = right[j]
                j += 1
            k += 1
        while i < len(left):
            array[k] = left[i]
            i += 1
            k += 1
        while j < len(right):
            array[k] = right[j]
            j += 1
            k += 1

