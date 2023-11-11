class Solution:
    def findLongestChain(self, pairs: List[List[int]]) -> int:
        print(pairs)
        self.mergesort(pairs, 0, len(pairs) - 1)
        print(pairs)
        return 0
    
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
            if left[i][0] < right[j][0]:
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

