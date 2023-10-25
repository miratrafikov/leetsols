func majorityElement(nums []int) int {
    mergesort(nums, 0, len(nums)-1)
    mid := (len(nums) - 1) / 2
    return nums[mid]
}

func mergesort(nums []int, l, h int) {
    if l == h {
        return
    }
    m := l + (h - l) / 2
    mergesort(nums, l, m)
    mergesort(nums, m+1, h)
    merge(nums, l, m, h)
}

func merge(nums []int, l, m, h int) {
    sortedNums, k := make([]int, h-l+1), 0
    i := 0
    j := 0
    for l+i <= m && m+1+j <= h {
        if nums[l+i] < nums[m+1+j] {
            sortedNums[k] = nums[l+i]
            i++
        } else {
            sortedNums[k] = nums[m+1+j]
            j++
        }
        k++
    }
    for l+i <= m {
        sortedNums[k] = nums[l+i]
        i++
        k++
    }
    for m+1+j <= h {
        sortedNums[k] = nums[m+1+j]
        j++
        k++
    }
    i = 0
    for i < len(sortedNums) {
        nums[l+i] = sortedNums[i]
        i++
    }
}
