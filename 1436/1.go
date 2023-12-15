func destCity(paths [][]string) string {
    var answer string
    for i := 0; i < len(paths); i++ {
        destinationCity := paths[i][1]
        isDepartureCity := false
        for j := 0; j < len(paths); j++ {
            if j != i && paths[j][0] == destinationCity {
                isDepartureCity = true
                break
            }
        }
        if !isDepartureCity {
            answer = destinationCity
            break
        }
    }
    return answer
}
