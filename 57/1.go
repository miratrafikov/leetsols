func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    result := make([][]int, 0, n+1)
    newA, newB := newInterval[0], newInterval[1]
    // append to result all intervals that end before new interval starts
    i := 0
    for i < n {
        if intervals[i][1] < newA {
            result = append(result, intervals[i])
            i++
        } else {
            break
        }
    }
    // append to result new interval merged with overlapping intervals if necessary
    a, b := newA, newB
    for i < n {
        if intervals[i][0] > b {
            break
        }
        a = min(a, intervals[i][0])
        b = max(b, intervals[i][1])
        i++
    }
    result = append(result, []int{a, b})
    // append to result the rest of the intervals
    for i < n {
        result = append(result, intervals[i])
        i++
    }
    return result
}
