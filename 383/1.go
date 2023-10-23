func canConstruct(ransomNote string, magazine string) bool {
    mapSrc := make(map[rune]int)
    for _, char := range magazine {
        if count, found := mapSrc[char]; found {
            mapSrc[char] = count + 1
        } else {
            mapSrc[char] = 1
        }
    }
    for _, char := range ransomNote {
        if count, found := mapSrc[char]; found {
            if count > 1 {
                mapSrc[char] = count - 1
            } else {
                delete(mapSrc, char)
            }
        } else {
            return false
        }
    }
    return true
}
