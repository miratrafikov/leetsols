func winnerOfGame(colors string) bool {
    aliceMovesCount, bobMovesCount :=  0, 0
    l, r := 0, 0
    for r < len(colors) {
        if colors[r] == colors[l] {
            r++
        }
        if r == len(colors) || colors[r] != colors[l] {
            sequenceLength := r - l
            if sequenceLength >= 3 {
                movesInSequence := sequenceLength - 2
                sequenceColor := colors[l]
                if sequenceColor == 'A' {
                    aliceMovesCount += movesInSequence
                } else {
                    bobMovesCount += movesInSequence
                }
            }
            l = r
        }
    }
    return aliceMovesCount > bobMovesCount   
}
