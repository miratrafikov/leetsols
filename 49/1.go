func groupAnagrams(strs []string) [][]string {
    descriptorLength := int('z') - int('a') + 1
    descriptorToAnagramGroup := make(map[string][]string)
    for _, s := range strs {
        lettersCountInS := make([]int, descriptorLength)
        for _, letter := range s {
            lettersCountInS[int(letter) - int('a')] += 1
        }
        descriptor := ""
        for i, count := range lettersCountInS {
            for j := 0; j < count; j++ {
                descriptor += string(rune(int('a') + i))
            }
        }
        descriptorToAnagramGroup[descriptor] = append(descriptorToAnagramGroup[descriptor], s)
    }
    groups := make([][]string, len(descriptorToAnagramGroup))
    i := 0
    for _, anagramGroup := range descriptorToAnagramGroup {
        groups[i] = anagramGroup
        i++
    }
    return groups
}
