const lettersInLanguage = 26

func partitionLabels(s string) []int {
    n := len(s)
    letterToPositions := make(map[byte][]int, lettersInLanguage)   
    for i := 0; i < n; i++ {
        positions, found := letterToPositions[s[i]]
        if !found {
            letterToPositions[s[i]] = []int{i,i}
        } else {
            letterToPositions[s[i]] = []int{positions[0],i}
        }
    }
    partitionLengths := []int{}
    partitionStart := 0
    for partitionStart < n {
        startingLetterPositions, _ := letterToPositions[s[partitionStart]]
        partitionEnd := startingLetterPositions[1]
        for i := partitionStart+1; i < partitionEnd; i++ {
            letterPositions, _ := letterToPositions[s[i]]
            if letterPositions[1] > partitionEnd {
                partitionEnd = letterPositions[1]
            }
        }
        partitionLengths = append(partitionLengths, partitionEnd - partitionStart + 1)
        partitionStart = partitionEnd + 1
    }
    return partitionLengths
}
