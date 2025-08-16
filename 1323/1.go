func maximum69Number (num int) int {
    // rd contains digits of number in reverse order
    rd := []byte{}
    numTemp := num
    for numTemp != 0 {
        tail := numTemp % 10
        numTemp = numTemp / 10
        if tail == 6 {
            rd = append(rd, '6')
        } else {
            rd = append(rd, '9')
        }
    }
    for i := len(rd) - 1; i >= 0; i-- {
        if rd[i] == '6' {
            rd[i] = '9'
            return rdToNum(rd)
        }
    }
    return num
}

func rdToNum(rd []byte) int {
    num := 0
    for i := len(rd) - 1; i >= 0; i-- {
        num *= 10
        var tail int
        if rd[i] == '6' {
            tail = 6
        } else {
            tail = 9
        }
        num += tail
    }
    return num
}
