func merge(intervals [][]int) [][]int {
    n := len(intervals)
    mergesort(intervals, 0, n-1)
    merged := [][]int{}
    l := 0
    for l < n {
        r := l
        maxRight := intervals[r][1]
        for r < n - 1 && maxRight >= intervals[r+1][0] {
            r++
            if intervals[r][1] > maxRight {
                maxRight = intervals[r][1]
            }
        }
        merged = append(merged, []int{intervals[l][0], maxRight})
        l = r+1
    }
    return merged
}

func mergesort(intervals [][]int, l, r int) {
    if l == r {
        return
    }
    midpoint := l + (r - l) / 2
    mergesort(intervals, l, midpoint)
    mergesort(intervals, midpoint+1, r)
    mergePartitions(intervals, l, midpoint, r)
}

func mergePartitions(intervals [][]int, l, midpoint, r int) {
    merged := make([][]int, 0, r-l+1)
    i := l
    j := midpoint+1
    for i <= midpoint && j <= r {
        if intervals[i][0] < intervals[j][0] {
            merged = append(merged, intervals[i])
            i++
        } else {
            merged = append(merged, intervals[j])
            j++
        }
    }
    for i <= midpoint {
        merged = append(merged, intervals[i])
        i++
    }
    for j <= r {
        merged = append(merged, intervals[j])
        j++
    }
    for i := 0; i < r-l+1; i++ {
        intervals[l + i] = merged[i]
    }
}
