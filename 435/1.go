func eraseOverlapIntervals(intervals [][]int) int {
    sortIntervalsByRightBoundary(intervals)
    maxNumberOfNonOverlappingIntervals := 1
    rightBoundaryOfMostRecentInterval := intervals[0][1]
    for _, interval := range intervals[1:] {
        if interval[0] >= rightBoundaryOfMostRecentInterval {
            maxNumberOfNonOverlappingIntervals++
            rightBoundaryOfMostRecentInterval = interval[1]
        }
    }
    return len(intervals) - maxNumberOfNonOverlappingIntervals
}

func sortIntervalsByRightBoundary(intervals [][]int) {
    quicksort(intervals, 0, len(intervals)-1)
}

func quicksort(intervals [][]int, l, r int) {
    if l == r {
        return
    }
    m := partition(intervals, l, r)
    quicksort(intervals, l, m)
    quicksort(intervals, m+1, r)
}

func partition(intervals [][]int, l, r int) int {
    pivot := intervals[l + (r - l) / 2]
    i := l - 1
    j := r + 1
    for {
        for {
            i++
            if greaterOrEqual(intervals[i], pivot) {
                break
            }
        }
        for {
            j--
            if lesserOrEqual(intervals[j], pivot) {
                break
            }
        }
        if i >= j {
            break
        }
        intervals[i], intervals[j] = intervals[j], intervals[i]
    }
    return j
}

func greaterOrEqual(a, b []int) bool {
    return a[1] >= b[1]
}

func lesserOrEqual(a, b []int) bool {
    return a[1] <= b[1]
}
