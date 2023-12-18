func divideArray(nums []int, k int) [][]int {
    n := len(nums)
    mergesort(nums, 0, n-1)
    numberOfGroups := n / 3
    groups := make([][]int, numberOfGroups)
    for i := 0; i < numberOfGroups; i++ {
        group := make([]int, 3)
        group[0] = nums[i * 3]
        group[1] = nums[i * 3 + 1]
        group[2] = nums[i * 3 + 2]
        if group[2] - group[0] > k {
            return [][]int{}
        }
        groups[i] = group
    }
    return groups
}

func mergesort(nums []int, l, r int) {
    if l == r {
        return
    }
    m := l + (r - l) / 2
    mergesort(nums, l, m)
    mergesort(nums, m+1, r)
    merge(nums, l, m ,r)
}

func merge(nums []int, l, m, r int) {
    sorted := make([]int, r-l+1)
    k := 0
    i := l
    j := m + 1
    for i <= m && j <= r {
        var number int
        if less(nums[i], nums[j]) {
            number = nums[i]
            i++
        } else {
            number = nums[j]
            j++
        }
        sorted[k] = number
        k++
    }
    for i <= m {
        sorted[k] = nums[i]
        i++
        k++
    }
    for j <= r {
        sorted[k] = nums[j]
        j++
        k++
    }
    for k := 0; k < len(sorted); k++ {
        nums[l + k] = sorted[k]
    }
}

func less(a, b int) bool {
    return a < b
}
