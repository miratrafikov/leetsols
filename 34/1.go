func searchRange(nums []int, target int) []int {
    n := len(nums)
    if n == 0 {
        return []int{-1,-1}
    }
    l := findLeftBoundary(nums, target, 0, n-1)
    r := findRightBoundary(nums, target, 0, n-1)
    return []int{l,r}
}

func findLeftBoundary(nums []int, target, l, r int) int {
    if l > r {
        return -1
    }
    m := l + (r - l) / 2
    if nums[m] < target {
        return findLeftBoundary(nums, target, m+1, r)
    }
    if nums[m] > target {
        return findLeftBoundary(nums, target, l, m-1)
    }
    resultFromLeft := findLeftBoundary(nums, target, l, m-1)
    if resultFromLeft != -1 {
        return resultFromLeft
    }
    return m
}

func findRightBoundary(nums []int, target, l, r int) int {
    if l > r {
        return -1
    }
    m := l + (r - l) / 2
    if nums[m] < target {
        return findRightBoundary(nums, target, m+1, r)
    }
    if nums[m] > target {
        return findRightBoundary(nums, target, l, m-1)
    }
    resultFromRight := findRightBoundary(nums, target, m+1, r)
    if resultFromRight != -1 {
        return resultFromRight
    }
    return m
}
