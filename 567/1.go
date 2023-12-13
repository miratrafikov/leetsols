// pattern: s1
// text: s2
func checkInclusion(s1 string, s2 string) bool {
    alphabetLength := 'z' - 'a' + 1
    patternDescriptor := make([]int, alphabetLength)
    for i := 0; i < len(s1); i++ {
        patternDescriptor[getIdOfChar(s1[i])]++
    }
    textDescriptor := make([]int, alphabetLength)
    for i := 0; i < len(s2); i++ {
        if i >= len(s1) {
            if areSlicesEqual(patternDescriptor, textDescriptor) {
                return true
            }
            textDescriptor[getIdOfChar(s2[i - len(s1)])]--
        }
        textDescriptor[getIdOfChar(s2[i])]++
    }
    if areSlicesEqual(patternDescriptor, textDescriptor) {
        return true
    }
    return false
}

func getIdOfChar(char byte) int {
    asciiCode := int(char)
    asciiOffset := int('a')
    return asciiCode - asciiOffset
}

func areSlicesEqual(s1, s2 []int) bool {
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            return false
        }
    }
    return true
}
