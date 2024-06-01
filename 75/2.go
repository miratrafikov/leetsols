func sortColors(nums []int)  {
    n := len(nums)
    // the loop below grows a cluster of 0s at the beginning of nums and a cluster of 2s at the 
    // end of nums
    positionAfter0s := 0
    positionBefore2s := n-1
    // iterate over the loop until we hit the start of cluster of 2s (we don't need to go beyond the start
    // of the cluster since we know that the cluster consists of 2s exclusively)
    i := 0
    for i <= positionBefore2s {
        // we strive for the following loop invariant at the end of each iteration:
        //  - there is a cluster of 0s that ends at positionAfter0s-1
        //  - there is a cluster of 2s that starts at positionBefore2s+1
        //  - there is a cluster of 1s that starts at positionAfter0s and ends at i-1
        //  - there is an unvisited portion starting at i and ending at positionBefore2s
        //  - 0000|1|02102|222222    
        switch nums[i] {
        case 2:
            nums[positionBefore2s], nums[i] = nums[i], nums[positionBefore2s]
            positionBefore2s--
        case 0:
            nums[positionAfter0s], nums[i] = nums[i], nums[positionAfter0s]
            positionAfter0s++
            i++
        case 1:
            i++
        }
    }
}
