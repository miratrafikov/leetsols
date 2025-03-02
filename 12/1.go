type romanAndDecimal struct {
	Roman   string
	Decimal int
}

var romans = []romanAndDecimal{
    {
        Roman: "I",
        Decimal: 1,
    },
    {
        Roman: "V",
        Decimal: 5,
    },
    {
        Roman: "X",
        Decimal: 10,
    },
    {
        Roman: "L",
        Decimal: 50,
    },
    {
        Roman: "C",
        Decimal: 100,
    },
    {
        Roman: "D",
        Decimal: 500,
    },
    {
        Roman: "M",
        Decimal: 1000,
    },
}

func intToRoman(num int) string {
	r := ""
	for num > 0 {
		leftmostDigit := getLeftmostDigit(num)
        if leftmostDigit != 4 && leftmostDigit != 9 {
            i := len(romans)-1
            for i >= 0 {
                if num >= romans[i].Decimal {
                    break
                }
                i--
            }
            r += romans[i].Roman
            num -= romans[i].Decimal
        } else {
            i := 0
            for num > romans[i].Decimal {
                i++
            }
            sumToSubtractFromNum := romans[i].Decimal
            rightRomanDigitOfPair := romans[i].Roman
            var leftRomanDigitOfPair string
            if leftmostDigit == 9 {
                leftRomanDigitOfPair = romans[i-2].Roman
                sumToSubtractFromNum -= romans[i-2].Decimal
            } else {
                leftRomanDigitOfPair = romans[i-1].Roman
                sumToSubtractFromNum -= romans[i-1].Decimal
            }
            r += leftRomanDigitOfPair + rightRomanDigitOfPair
            num -= sumToSubtractFromNum
        }
	}
	return r
}

func getLeftmostDigit(num int) int {
	for num >= 10 {
		num /= 10
	}
	return num
}
