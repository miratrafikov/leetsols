from typing import List

def merge(nums: List[int], p: int, q: int, r: int):
    nums_left_len = q - p + 1
    nums_right_len = r - q
    nums_left = [None] * nums_left_len
    nums_right = [None] * nums_right_len

    for i in range(nums_left_len):
        nums_left[i] = nums[p + i]
    
    for j in range(nums_right_len):
        nums_right[j] = nums[q + 1 + j]
    
    i = j = 0
    k = p
    
    while i < nums_left_len and j < nums_right_len:
        if nums_left[i] < nums_right[j]:
            nums[k] = nums_left[i]
            i += 1
        else:
            nums[k] = nums_right[j]
            j += 1
        
        k += 1
    
    while i < nums_left_len:
        nums[k] = nums_left[i]
        i += 1
        k += 1
    
    while j < nums_right_len:
        nums[k] = nums_right[j]
        j += 1
        k += 1

def mergesort(nums: List[int], p: int, r: int):
    if r - p == 0:
        return
    
    q = p + (r - p) // 2
    mergesort(nums, p, q)
    mergesort(nums, q + 1, r)
    merge(nums, p, q, r)
        

def sort(nums: List[int]):
    mergesort(nums, 0, len(nums) - 1)

arr = [4, 3, 5, 6, 2]
sort(arr)
print(arr)
