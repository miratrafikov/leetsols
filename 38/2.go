func countAndSay(n int) string {
    casPrev := "1"
    for i := 1; i < n; i++ {
        var sb strings.Builder
        for j := 0; j < len(casPrev); j++ {
            targetDigit := casPrev[j]
            k := 1
            for k+j < len(casPrev) {
                if casPrev[k+j] == targetDigit {
                    k++
                } else {
                    break
                }
            }
            j += k-1
            sb.Write([]byte(strconv.Itoa(k)))
            sb.Write([]byte{targetDigit})
        }
        casPrev = sb.String()
    }
    return casPrev
}

// 1 — 1
// 2 — 11
// 3 — 21
// 4 — 1211
// 5 - 111221
// 6 - 312211
// 7 - 13112221
