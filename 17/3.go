func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    digitToLetters := mapDigitToLetters()
    combinations, _ := digitToLetters[digits[0]]
    for _, digit := range digits[1:] {
        letters, _ := digitToLetters[byte(digit)]
        newCombinations := []string{}
        for _, letter := range letters {
            for _, combination := range combinations {
                newCombinations = append(newCombinations, combination + letter)
            }
        }
        combinations = newCombinations
    }
    return combinations
}

func mapDigitToLetters() map[byte][]string {
    return map[byte][]string{
        '2': []string{"a", "b", "c"},
        '3': []string{"d", "e", "f"},
        '4': []string{"g", "h", "i"},
        '5': []string{"j", "k", "l"},
        '6': []string{"m", "n", "o"},
        '7': []string{"p", "q", "r", "s"},
        '8': []string{"t", "u", "v"},
        '9': []string{"w", "x", "y", "z"},
    }
}
