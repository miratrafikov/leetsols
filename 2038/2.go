func winnerOfGame(colors string) bool {
    winningScale := 0
    for i := 1; i < len(colors) - 1; i++ {
        if colors[i - 1] == colors[i] && colors[i] == colors[i + 1] {
            if colors[i] == 'A' {
                winningScale--
            } else {
                winningScale++
            }
        }
    }
    return winningScale < 0
}
