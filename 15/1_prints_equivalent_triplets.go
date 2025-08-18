func threeSum(nums []int) [][]int {
    n := len(nums)
    valToPositionDuplets := map[int][][]int{}
    for i := 0; i < n-1; i++ {
        for j := i+1; j < n; j++ {
            val := nums[i] + nums[j]
            valToPositionDuplets[val] = append(valToPositionDuplets[val], []int{i,j})
        }
    }
    result := [][]int{}
    for i := 0; i < n-2; i++ {
        if duplets, ok := valToPositionDuplets[-nums[i]]; ok {
            for _, duplet := range duplets {
                if duplet[0] > i && duplet[1] > i {
                    result = append(result, []int{nums[i], nums[duplet[0]], nums[duplet[1]]})
                }
            }
        }
    }
    return result
}
