// we will use a matrix of the same shape
// that matrix will contain 0s in positions corresponding to the positions of 0s in the input matrix and -1 elsewhere
// then we will spawn BFS from the positions of zeroes and simulate processing them step by step in parallel
// each emptying of the current queue is one tick
// if an element has been visited, we may skip it as it was surely accessed earlier and therefore a shorter path was already found
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
    currentTick := 0
    currentTickElementsCount := len(q)
    for currentTickElementsCount > 0 {
        // take one element from the queue
        cell := q[0]
        q = q[1:]
        fmt.Println(q)
        fmt.Println(distances)
        currentTickElementsCount--
        if distances[cell[0]][cell[1]] == -1 {
            // the element's value is the amount of ticks that passed so far
            distances[cell[0]][cell[1]] = currentTick
        }
        // add surrounding elements
        addSurroudingCellsToQueue(cell, &q, distances, m, n)
        // if the tick is done, start a new one
        if currentTickElementsCount == 0 {
            currentTick++
            currentTickElementsCount = len(q)
        }
    }
    return distances
}

func addSurroudingCellsToQueue(cell []int, q *[][]int, distances [][]int, m, n int) {
    // calculate positions of 4 elements that surround the element
    top := []int{cell[0]-1, cell[1]}
    bottom := []int{cell[0]+1, cell[1]}
    left := []int{cell[0], cell[1]-1}
    right := []int{cell[0], cell[1]+1}
    // push them to a queue if they are not out of matrix and were not visited already
    if !isOutOfMatrix(top, m, n) && distances[top[0]][top[1]] == -1 {
        *q = append(*q, top)
    }
    if !isOutOfMatrix(bottom, m, n) && distances[bottom[0]][bottom[1]] == -1 {
        *q = append(*q, bottom)
    }
    if !isOutOfMatrix(left, m, n) && distances[left[0]][left[1]] == -1 {
        *q = append(*q, left)
    }
    if !isOutOfMatrix(right, m, n) && distances[right[0]][right[1]] == -1 {
        *q = append(*q, right)
    }
}

func isOutOfMatrix(cell []int, m, n int) bool {
    return cell[0] < 0 || cell[0] > m - 1 || cell[1] < 0 || cell[1] > n - 1
}
