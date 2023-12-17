func merge(intervals [][]int) [][]int {
    n := len(intervals)
    quicksort(intervals, 0, n-1)
    result := [][]int{}
    start, end := intervals[0][0], intervals[0][1]
    for i := 1; i < n; i++ {
        if intervals[i][0] <= end {
            if intervals[i][1] > end {
                end = intervals[i][1]
            }
        } else {
            result = append(result, []int{start, end})
            start, end = intervals[i][0], intervals[i][1]
        }
    }
    result = append(result, []int{start, end})
    return result
}

func quicksort(intervals [][]int, l, r int) {
    if l == r {
        return
    }
    pivotIndex := partition(intervals, l, r)
    quicksort(intervals, l, pivotIndex)
    quicksort(intervals, pivotIndex+1, r)
}

func partition(intervals [][]int, l, r int) int {
    pivot := intervals[l + (r - l) / 2]
    i := l - 1
    j := r + 1
    for {
        for {
            i++
            if !less(intervals[i], pivot) {
                break
            }
        }
        for {
            j--
            if !less(pivot, intervals[j]) {
                break
            }
        }
        if i >= j {
            break
        }
        swap(intervals, i, j)
    }
    return j
}

func less(a, b []int) bool {
    if a[0] < b[0] {
        return true
    }
    return false
}

func swap(intervals [][]int, i, j int) {
    intervals[i], intervals[j] = intervals[j], intervals[i]
}
