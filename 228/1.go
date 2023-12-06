func summaryRanges(nums []int) []string {
    ranges := []string{}
    if len(nums) == 0 {
        return ranges
    }
    rangeStart, rangeEnd := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] + 1 == nums[i] {
            rangeEnd = nums[i]
        } else {
            if rangeStart != rangeEnd {
                ranges = append(ranges, fmt.Sprintf("%d->%d", rangeStart, rangeEnd)) 
            } else {
                ranges = append(ranges, fmt.Sprintf("%d", rangeStart))
            }
            rangeStart, rangeEnd = nums[i], nums[i]
        }
    }
    if rangeStart != rangeEnd {
        ranges = append(ranges, fmt.Sprintf("%d->%d", rangeStart, rangeEnd)) 
    } else {
        ranges = append(ranges, fmt.Sprintf("%d", rangeStart))
    }
    return ranges
}
