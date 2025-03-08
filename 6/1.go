func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    result := ""
    n := len(s)
    for i := 0; i < numRows; i++ {
        steps := stepsForRow(i+1, numRows)
        isCurrentStepOdd := true
        j := i
        for j < n {
            result += string(s[j])
            if isCurrentStepOdd {
                j += steps[0]
            } else {
                j += steps[1]
            }
            isCurrentStepOdd = !isCurrentStepOdd
        }
    }
    return result
}

func stepsForRow(row, totalRows int) [2]int {
    verticalLength := totalRows - 1
    diagonalLength := totalRows - 1
    maxWalkInPattern := verticalLength + diagonalLength
    if row == 1 || row == totalRows  {
        return [2]int{maxWalkInPattern, maxWalkInPattern}
    }
    maxWalkInPatternForRow := maxWalkInPattern - 2 * (row - 1)
    return [2]int{maxWalkInPatternForRow, maxWalkInPattern - maxWalkInPatternForRow}
}
