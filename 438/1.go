func findAnagrams(s string, p string) []int {
    m, n := len(p), len(s)
    var result []int
    if n < m {
        return result
    }
    patFlags, txtFlags := make([]int16, 26), make([]int16, 26)
    for i := 0; i < m; i++ {
        patFlags[letterToIndex(p[i])]++
        txtFlags[letterToIndex(s[i])]++
    }
    if areEqual(patFlags, txtFlags) {
        result = append(result, 0)
    }
    for i := m; i < n; i++ {
        txtFlags[letterToIndex(s[i-m])]--
        txtFlags[letterToIndex(s[i])]++
        if areEqual(patFlags, txtFlags) {
            result = append(result, i-m+1)
        }
    }
    return result
}

func letterToIndex(letter byte) byte {
    return letter - 'a'
}

func areEqual(s1, s2 []int16) bool {
    equal := true
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            equal = false
            break
        }
    }
    return equal
}
