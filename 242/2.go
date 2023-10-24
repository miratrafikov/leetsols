func isAnagram(s string, t string) bool {
    m := make(map[rune]int)
    for _, char := range s {
        if count, found := m[char]; found {
            m[char] = count + 1
        } else {
            m[char] = 1
        }
    }
    for _, char := range t {
        if count, found := m[char]; found {
            if count == 0 {
                return false
            } else if count == 1 {
                delete(m, char)
            } else {
                m[char] = count - 1
            }
        } else {
            return false
        }
    }
    return len(m) == 0
}
