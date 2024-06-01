const alphabetSize = int('z') - int('A') + 1

func minWindow(s string, t string) string {
    n := len(s)
    m := len(t)
    // 52 is alphabet size
    var mapWindow [alphabetSize]int
    var mapTarget [alphabetSize]int
    for i := 0; i < m; i++ {
        mapTarget[letterToIndex(t[i])]++
    }
    i, j := 0, 0
    isAnyMatchFound := false
    resI, resJ := 0, 0
    for j < n {
        mapWindow[letterToIndex(s[j])]++
        for covers(mapWindow, mapTarget) {
            if !isAnyMatchFound || j - i < resJ - resI {
                resI, resJ = i, j
                isAnyMatchFound = true
            }
            mapWindow[letterToIndex(s[i])]--
            i++
        }
        j++
    }
    if !isAnyMatchFound {
        return ""
    }
    return s[resI:resJ+1]
}

func letterToIndex(letter byte) int {
    if int(letter) - int('A') == 54 {
        fmt.Println(rune(letter))
    }
    return int(letter) - int('A')
}

func covers(m1, m2 [alphabetSize]int) bool {
    for i := 0; i < alphabetSize; i++ {
        if m1[i] < m2[i] {
            return false
        }
    }
    return true
}
