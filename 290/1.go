func wordPattern(pattern string, s string) bool {
    words := strings.Fields(s)
    if len(words) != len(pattern) {
        return false
    }
    letterToWord := make(map[rune]string)
    usedWords := make(map[string]bool)
    for i, letter := range pattern {
        if word, found := letterToWord[letter]; found {
            if word != words[i] {
                return false
            }
        } else {
            if _, found := usedWords[words[i]]; found {
                return false
            }
            letterToWord[letter] = words[i]
            usedWords[words[i]] = true
        }
    }
    return true
}
