func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    // startLocations nodes contain either 0, -1, or 1.
    // If a node is a suitable start, it will be 0.
    // If a node was a suitable start and is being test or was tested for match, it will be 1.
    // If a node is not a suitable start, it will be -1.
    startLocations := makeMatrix[int](m, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] != word[0] {
                startLocations[i][j] = -1
            }
        }
    }
    // Test every suitable start location.
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if startLocations[i][j] == 0 {
                startLocations[i][j] = 1
                if testStartLocation(i, j, board, word) {
                    return true
                }
            }
        }
    }
    return false
}

func makeMatrix[T int | bool](m, n int) [][]T {
    matrix := make([][]T, m)
    for i := 0; i < m; i++ {
        matrix[i] = make([]T, n)
    }
    return matrix
}

type position struct {
    I, J int
}

func testStartLocation(i, j int, board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    letter := 0
    // visited contains nodes that are false by default and then marked as true if they were visited.
    visited := makeMatrix[bool](m, n)
    return dfs(i, j, board, word, letter, visited)
}

func dfs(i, j int, board [][]byte, word string, letter int, visited [][]bool) bool {
    if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || visited[i][j] {
        return false
    }
    if board[i][j] != word[letter] {
        return false
    }
    visitedCopy := make([][]bool, len(visited))
    for k := range visited {
        visitedCopy[k] = make([]bool, len(visited[0]))
        copy(visitedCopy[k], visited[k])
    }
    visitedCopy[i][j] = true
    letter++
    if letter == len(word) {
        return true
    }
    return dfs(i-1, j, board, word, letter, visitedCopy) ||
        dfs(i+1, j, board, word, letter, visitedCopy) ||
        dfs(i, j-1, board, word, letter, visitedCopy) ||
        dfs(i, j+1, board, word, letter, visitedCopy)
}

