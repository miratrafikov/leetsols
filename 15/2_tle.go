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
                if i == duplet[0] || i == duplet[1] {
                    continue
                }
                tripletSorted := sortTriplet(nums[i], nums[duplet[0]], nums[duplet[1]])
                if !contains(result, tripletSorted) {
                    result = append(result, tripletSorted)
                }
            }
        }
    }
    return result
}

func sortTriplet(a, b, c int) []int {
    s := []int{a, b, c}
    if s[0] > s[1] {
        s[0], s[1] = s[1], s[0]
    }
    if s[1] > s[2] {
        s[1], s[2] = s[2], s[1]
    }
    if s[0] > s[1] {
        s[0], s[1] = s[1], s[0]
    }
    return s
}

func contains(arr [][]int, target []int) bool {
    for i := 0; i < len(arr); i++ {
        if arr[i][0] == target[0] && arr[i][1] == target[1] && arr[i][2] == target[2] {
            return true
        }
    }
    return false
}
