func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := mergeSlices(nums1, nums2)
    n := len(merged)
    if n % 2 != 0 {
        return float64(merged[n / 2])
    }
    return float64(merged[n / 2 - 1] + merged[n / 2]) / 2.0
}

func mergeSlices(s1, s2 []int) []int {
    m, n := len(s1), len(s2)
    merged := make([]int, m+n)
    i, j, k := 0, 0, 0
    for i < m && j < n {
        if s1[i] < s2[j] {
            merged[k] = s1[i]
            i++
        } else {
            merged[k] = s2[j]
            j++
        }
        k++
    }
    for i < m {
        merged[k] = s1[i]
        i++
        k++
    }
    for j < n {
        merged[k] = s2[j]
        j++
        k++
    }
    return merged
}
