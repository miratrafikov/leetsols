import "math"
import "strconv"

func intFromBinaryString(s string) int {
    x := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            exponent := len(s) - i - 1
            x += int(math.Pow(float64(2), float64(exponent)))
        }
    }

    return x
}

func intToBinaryString(x int) string {
    if x == 0 {
        return "0"
    }

    digits := []int{}
    for x > 0 {
        remainder := x % 2
        x /= 2
        digits = append(digits, remainder)
    }

    s := ""
    for i := 0; i < len(digits); i++ {
        digit := digits[len(digits) - 1 - i]
        s = s + strconv.Itoa(digit)
    }

    return s
}

func addBinary(a string, b string) string {
    x := intFromBinaryString(a)
    y := intFromBinaryString(b)
    z := x + y

    return intToBinaryString(z)
}
