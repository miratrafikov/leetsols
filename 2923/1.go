func findChampion(grid [][]int) int {
    n := len(grid)
    var teamThatNeverLost int
    for teamToTestIfEverLost := 0; teamToTestIfEverLost < n; teamToTestIfEverLost++ {
        didLose := false
        for rivalTeam := 0; rivalTeam < n; rivalTeam++ {
            if teamToTestIfEverLost != rivalTeam && grid [rivalTeam][teamToTestIfEverLost] == 1 {
                didLose = true
                break
            }
        }
        if !didLose {
            teamThatNeverLost = teamToTestIfEverLost
            break
        }
    }
    return teamThatNeverLost
}
