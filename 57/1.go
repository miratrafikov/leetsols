func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    result := make([][]int, 0, n+1)
    // append to result all intervals that end before new interval starts
    i := 0
    for i < n {
        if intervals[i][1] < newInterval[0] {
            result = append(result, intervals[i])
            i++
        } else {
            break
        }
    }
    // we are now either at the end of the list, so all intervals in the list were "lesser" than 
    // the new interval, or need to merge the new interval with the current interval (pointed at by i).
    // If case 1, just append new interval to result.
    // If case 2, we make an interval from merging the new interval and the interval at i and keep doing that
    // until we are either at the end of the list of found an interval that starts after the merge-interval
    // ends.
    mergedInterval := newInterval
    for i < n {
        if intervals[i][0] > mergedInterval[1] {
            break
        }
        mergedInterval[0] = min(mergedInterval[0], intervals[i][0])
        mergedInterval[1] = max(mergedInterval[1], intervals[i][1])
        i++
    }
    result = append(result, mergedInterval)
    // append to result the rest of the intervals
    for i < n {
        result = append(result, intervals[i])
        i++
    }
    return result
}
