func maxDistinctElements(nums []int, k int) int {
    // quicksort(nums, 0, len(nums)-1)
    sort.Ints(nums)
    // used stores distinct values.
    used := make(map[int]bool)
    // If key is a number from nums, mins[key] is the largest tried 
    // k-modificator as long as we try k-modificator from -k to k.
    mins := make(map[int]int)
    for _, num := range nums {
        init := -k
        if minkmod, found := mins[num]; found {
            init = minkmod+1
        }
        for kmod := init; kmod <= k; kmod++ {
            target := num + kmod
            if !used[target] {
                used[target] = true
                mins[num] = kmod
                break
            }
        }
    }
    return len(used)
}

func quicksort(nums []int, l, r int) {
    if l >= r {
        return
    }
    p := partition(nums, l ,r)
    quicksort(nums, l, p)
    quicksort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
    pivot := nums[l + (r - l) / 2]
    i := l-1
    j := r+1
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
