func generateParenthesis(n int) []string {
    pairs := []string{}
    generate(&pairs, n, 1, "(")
    return pairs
}

func generate(pairs *[]string, pairsNeeded, openBracketsCount int, previousCombination string) {
    if len(previousCombination) == pairsNeeded * 2 {
        *pairs = append(*pairs, previousCombination)
        return
    }
    if openBracketsCount < pairsNeeded {
        generate(pairs, pairsNeeded, openBracketsCount+1, previousCombination + "(")
    }
    closedBracketsCount := len(previousCombination) - openBracketsCount
    if openBracketsCount > closedBracketsCount {
        generate(pairs, pairsNeeded, openBracketsCount, previousCombination + ")")
    }
}
