func findChampion(grid [][]int) int {
    n := len(grid)
    var teamThatNeverLost int
    for teamToTestIfItEverLost := 0; teamToTestIfItEverLost < n; teamToTestIfItEverLost++ {
        didLose := false
        for rivalTeam := 0; rivalTeam < n; rivalTeam++ {
            if grid [rivalTeam][teamToTestIfItEverLost] == 1 {
                didLose = true
                break
            }
        }
        if !didLose {
            teamThatNeverLost = teamToTestIfItEverLost
            break
        }
    }
    return teamThatNeverLost
}
