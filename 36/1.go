func isValidSudoku(board [][]byte) bool {
    for r := 0; r < 9; r++ {
        numsInRow := make(map[byte]struct{}, 9)
        for c := 0; c < 9; c++ {
            e := board[r][c]
            if e == '.' {
                continue
            }
            if _, ok := numsInRow[e]; ok {
                return false
            }
            numsInRow[e] = struct{}{}
        }
    }
    for c := 0; c < 9; c++ {
        numsInCol := make(map[byte]struct{}, 9)
        for r := 0; r < 9; r++ {
            e := board[r][c]
            if e == '.' {
                continue
            }
            if _, ok := numsInCol[e]; ok {
                return false
            }
            numsInCol[e] = struct{}{}
        }
    }
    for boxRow := 0; boxRow < 3; boxRow++ {
        for boxCol := 0; boxCol < 3; boxCol++ {
            numsInBox := make(map[byte]struct{}, 9)
            for r := 0; r < 3; r++ {
                for c := 0; c < 3; c++ {
                    eRow := boxRow * 3 + r
                    eCol := boxCol * 3 + c
                    e := board[eRow][eCol]
                    if e == '.' {
                        continue
                    }
                    if _, ok := numsInBox[e]; ok {
                        return false
                    }
                    numsInBox[e] = struct{}{}
                }
            }
        }
    }
    return true
}

