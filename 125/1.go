import (
    "regexp"
    "strings"
)

func isPalindrome(s string) bool {
    runes := []rune(s)
    left, right := 0, len(runes) - 1
    for left <= right {
        isLeftAlphanumeric, _ := regexp.MatchString("[a-zA-Z0-9]", string(runes[left]))
        if !isLeftAlphanumeric {
            left++
            continue
        }
        isRightAlphanumeric, _ := regexp.MatchString("[a-zA-Z0-9]", string(runes[right]))
        if !isRightAlphanumeric {
            right--
            continue
        }
        if strings.ToLower(string(runes[left])) != strings.ToLower(string(runes[right])) {
            return false
        }
        left++
        right--
    }
    return true
}
