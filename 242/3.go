func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    charToCount := make([]int, 'z'-'a'+1)
    asciiOffset := byte('a')
    for i := 0; i < len(s); i++ {
        charS, charT := int(s[i]-asciiOffset), int(t[i]-asciiOffset)
        charToCount[charS]++
        charToCount[charT]--
    }
    for i := 0; i < len(charToCount); i++ {
        if charToCount[i] != 0 {
            return false
        }
    }
    return true
}
