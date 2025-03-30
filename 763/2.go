const (
    lettersInLanguage = 26
    a = 97
)

func partitionLabels(s string) []int {
    n := len(s)
    lastOccurences := make([]int, lettersInLanguage)
    for i := 0; i < n; i++ {
        lastOccurences[s[i]-a] = i
    }
    lenghts := []int{}
    partitionStart := 0
    for partitionStart < n {
        partitionEnd := lastOccurences[s[partitionStart]-a]
        for i := partitionStart + 1; i < partitionEnd; i++ {
            lastOccurence := lastOccurences[s[i]-a]
            if lastOccurence > partitionEnd {
                partitionEnd = lastOccurence
            }
        }
        lenghts = append(lenghts, partitionEnd - partitionStart + 1)
        partitionStart = partitionEnd + 1
    }
    return lenghts
}
