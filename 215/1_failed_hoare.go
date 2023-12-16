func findKthLargest(nums []int, k int) int {
    // hehe := []int{5,4,7,12,365,21,1,4,2,66}
    return quickSelect(nums, k, 0, len(nums)-1)
}

func quickSelect(nums []int, k, l, h int) int {
    if l == h {
        return nums[l]
    }
    pivotIndex := partition(nums, l, h)
    if k == pivotIndex {
        return nums[pivotIndex]
    }
    if k < pivotIndex {
        return quickSelect(nums, k, l, pivotIndex-1)
    }
    return quickSelect(nums, k, pivotIndex+1, h)
}

func partition(nums []int, l, h int) int {
    pivot := nums[l + (h - l) / 2]
    i := l - 1
    j := h + 1
    for {
        for {
            i++
            if nums[i] >= pivot {
                break
            }
        }
        for {
            j--
            if nums[j] <= pivot {
                break
            }
        }
        if i >= j {
            break
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    return j
}
