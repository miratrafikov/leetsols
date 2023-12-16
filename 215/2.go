func findKthLargest(nums []int, k int) int {
    n := len(nums)
    l := 0
    r := n - 1
    expectedPosition := n - k
    return quickselect(nums, l, r, expectedPosition)
}

func quickselect(nums []int, l, r, k int) int {
    if l == r {
        return nums[l]
    }
    m := partition(nums, l, r)
    if m >= k {
        return quickselect(nums, l, m, k)
    } else {
        return quickselect(nums, m+1, r, k)
    }
}

func partition(nums []int, l, r int) int {
    pivot := nums[l + (r - l) / 2]
    i := l - 1
    j := r + 1
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
