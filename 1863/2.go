func subsetXORSum(nums []int) int {
    n := len(nums)
    subsetsXORs := make([]int, 0, int(math.Pow(2, float64(n))))
    xorAllSubsets(nums, 0, 0, &subsetsXORs)
    sumOfXORs := 0
    for i := 0; i < len(subsetsXORs); i++ {
        sumOfXORs += subsetsXORs[i]
    }
    return sumOfXORs
}

func xorAllSubsets(source []int, index, currentSubsetXOR int, subsetsXORs *[]int)  {
    if index == len(source) {
        *subsetsXORs = append(*subsetsXORs, currentSubsetXOR)
        return
    }
    xorAllSubsets(source, index+1, currentSubsetXOR ^ source[index], subsetsXORs)
    xorAllSubsets(source, index+1, currentSubsetXOR, subsetsXORs)
}

func copySlice(s []int) []int {
    dest := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        dest[i] = s[i]
    }
    return dest
}
