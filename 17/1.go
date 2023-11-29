func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    return appendNextDigitLetters(
        "",
        digits,
        0,
    )
}

func appendNextDigitLetters(
    currentCombination string,
    digits string,
    currentDigit int,
) []string {
    if currentDigit == len(digits) {
        return []string{currentCombination}
    }
    currentDigitLetters := lettersFromDigit(digits[currentDigit])
    currentCombinationAppendVariants := make([]string, len(currentDigitLetters))
    for i, letter := range currentDigitLetters {
        currentCombinationAppendVariants[i] = currentCombination + letter
    }
    nextVariants := make([]string, 0)
    for _, variant := range currentCombinationAppendVariants {
        nextVariants = append(nextVariants, appendNextDigitLetters(variant, digits, currentDigit+1)...)
    }
    return nextVariants
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
