func isAnagram(s string, t string) bool {
    mapS := make(map[rune]int)
    mapT := make(map[rune]int)
    for _, char := range s {
        if count, found := mapS[char]; found {
            mapS[char] = count + 1
        } else {
            mapS[char] = 1
        }
    }
    for _, char := range t {
        if count, found := mapT[char]; found {
            mapT[char] = count + 1
        } else {
            mapT[char] = 1
        }
    }
    if len(mapS) != len(mapT) {
        return false
    }
    for char, countS := range mapS {
        if countT, found := mapT[char]; !found || countT != countS {
            return false
        }
    }
    return true
}
