// lomuto
func findKthLargest(nums []int, k int) int {
    return quickselect(nums, 0, len(nums)-1, k-1)
}

func quickselect(nums []int, l, r, k int) int {
    if l == r {
        return nums[l]
    }
    pivotIndex := partition(nums, l, r)
    if pivotIndex > k {
        return quickselect(nums, l, pivotIndex-1, k)
    } else if pivotIndex < k {
        return quickselect(nums, pivotIndex+1, r, k)
    }
    return nums[pivotIndex]
}

func partition(nums []int, l, r int) int {
    pivotIndex := l + (r - l) / 2
    pivot := nums[pivotIndex]
    nums[pivotIndex], nums[r] = nums[r], nums[pivotIndex]
    i := l - 1
    for j := l; j < r; j++ {
        if nums[j] > pivot {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    i++
    nums[i], nums[r] = nums[r], nums[i]
    return i
}
