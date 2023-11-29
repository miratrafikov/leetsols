func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    combinations := lettersFromDigit(digits[0])
    for _, digit := range digits[1:] {
        letters := lettersFromDigit(byte(digit))
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

func lettersFromDigit(digit byte) []string {
    switch digit {
    case '2':
        return []string{"a", "b", "c"}
    case '3':
        return []string{"d", "e", "f"}
    case '4':
        return []string{"g", "h", "i"}
    case '5':
        return []string{"j", "k", "l"}
    case '6':
        return []string{"m", "n", "o"}
    case '7':
        return []string{"p", "q", "r", "s"}
    case '8':
        return []string{"t", "u", "v"}
    default:
        return []string{"w", "x", "y", "z"}
    }
}
