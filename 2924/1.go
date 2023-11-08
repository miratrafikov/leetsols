func findChampion(n int, edges [][]int) int {
    teamToWhetherItLost := make([]bool, n)
    for _, edge := range edges {
        teamToWhetherItLost[edge[1]] = true
    }
    teamThatNeverLost := -1
    for team, didLose := range teamToWhetherItLost {
        if !didLose {
            if teamThatNeverLost != -1 {
                return -1
            }
            teamThatNeverLost = team
        }
    }
    return teamThatNeverLost
}
