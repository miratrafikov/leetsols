func productExceptSelf(nums []int) []int {
	prefixes, suffixes := getPrefixesAndSuffixesArrays(nums)
    return getResultArray(nums, prefixes, suffixes)
}

func getPrefixesAndSuffixesArrays(nums []int) ([]int, []int) {
    n := len(nums)
    prefixes := make([]int, n)
	prefixes[0] = nums[0]
	suffixes := make([]int, n)
	suffixes[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		prefixes[i] = prefixes[i-1] * nums[i]
		suffixes[n-i-1] = suffixes[n-i] * nums[n-i-1]
	}
    return prefixes, suffixes
}

func getResultArray(nums, prefixes, suffixes []int) []int {
    n := len(nums)
    result := make([]int, n)
	result[0] = suffixes[1]
	result[n-1] = prefixes[n-2]
	for i := 1; i < n-1; i++ {
		result[i] = prefixes[i-1] * suffixes[i+1]
	}
	return result
}

