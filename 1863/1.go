func subsetXORSum(nums []int) int {
    n := len(nums)
    subsets := make([][]int, 0, int(math.Pow(2, float64(n))))
    generateSubsets(nums, 0, []int{}, &subsets)
    sumOfXORs := 0
    for i := 0; i < len(subsets); i++ {
        xor := 0
        for j := 0; j < len(subsets[i]); j++ {
            xor ^= subsets[i][j]
        }
        sumOfXORs += xor
    }
    return sumOfXORs
}

func generateSubsets(source []int, index int, acc []int, totalAcc *[][]int)  {
    if index == len(source) {
        *totalAcc = append(*totalAcc, acc)
        return
    }
    accCopy := copySlice(acc)
    generateSubsets(source, index+1, accCopy, totalAcc)
    acc = append(acc, source[index])
    generateSubsets(source, index+1, acc, totalAcc)
}

func copySlice(s []int) []int {
    dest := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        dest[i] = s[i]
    }
    return dest
}
