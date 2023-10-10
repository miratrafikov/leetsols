func updateMatrix(mat [][]int) [][]int {
    // get dimensions
    m := len(mat)
    n := len(mat[0])
    // create distances matrix and queue that we will use for the algorithm
    distances := [][]int{}
    q := [][]int{}
    // populate distances matrix after which we will no longer need the original matrix
    for i := 0; i < m; i++ {
        distances = append(distances, []int{})
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                q = append(q, []int{i, j})
                distances[i] = append(distances[i], 0)
            } else {
                distances[i] = append(distances[i], -1)
            }
        }
    }
    currentTickElementsCount := len(q)
    for currentTickElementsCount > 0 {
        // take one cell from the queue
        currentCell := q[0]
        q = q[1:]
        currentTickElementsCount--
        // get positions of the adjacent cell
        top := []int{currentCell[0]-1, currentCell[1]}
        bottom := []int{currentCell[0]+1, currentCell[1]}
        left := []int{currentCell[0], currentCell[1]-1}
        right := []int{currentCell[0], currentCell[1]+1}
        // check if each position is valid. If it is, check if it was not visited yet
        // If it wasn't, write new value to the cell and add it to the queue to visit in next queue tick
        distanceOfCurrentCell := distances[currentCell[0]][currentCell[1]]
        if !isOutOfMatrix(top, m, n) && distances[top[0]][top[1]] == -1 {
            distances[top[0]][top[1]] = distanceOfCurrentCell + 1
            q = append(q, top)
        }
        if !isOutOfMatrix(bottom, m, n) && distances[bottom[0]][bottom[1]] == -1 {
            distances[bottom[0]][bottom[1]] = distanceOfCurrentCell + 1
            q = append(q, bottom)
        }
        if !isOutOfMatrix(left, m, n) && distances[left[0]][left[1]] == -1 {
            distances[left[0]][left[1]] = distanceOfCurrentCell + 1
            q = append(q, left)
        }
        if !isOutOfMatrix(right, m, n) && distances[right[0]][right[1]] == -1 {
            distances[right[0]][right[1]] = distanceOfCurrentCell + 1
            q = append(q, right)
        }
        // if the current tick is done, start a new one
        if currentTickElementsCount == 0 {
            currentTickElementsCount = len(q)
        }
    }
    return distances
}

func isOutOfMatrix(cell []int, m, n int) bool {
    return cell[0] < 0 || cell[0] > m - 1 || cell[1] < 0 || cell[1] > n - 1
}
