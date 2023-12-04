import "strings"

func largestGoodInteger(num string) string {
    for digit := '9'; digit >= '0'; digit-- {
        runeTriplet := string([]rune{digit, digit, digit})
        if strings.Contains(num, runeTriplet) {
            return runeTriplet
        }
    }
    return ""
}
