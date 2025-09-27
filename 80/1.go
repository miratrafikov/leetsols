func removeDuplicates(nums []int) int {
    n := len(nums)

    swapBackDistance := 0
    memNum, memNumCount := nums[0], 0
    for i := 0; i < n; i++ {
        currentNum := nums[i]
        swap(nums, i, i-swapBackDistance)
        if currentNum == memNum {
            memNumCount++
            if memNumCount > 2 {
                swapBackDistance++
            }
        } else {
            memNum = currentNum
            memNumCount = 1
        }
    }

    return n - swapBackDistance
}

func swap(arr []int, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}
