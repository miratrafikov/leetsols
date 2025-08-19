type queue [][2]int

func newQueue() *queue {
    q := queue([][2]int{})
    return &q
}

func (q *queue) Enqueue(v [2]int) {
    *q = append(*q, v)
}

func (q *queue) Dequeue() [2]int {
    v := (*q)[0]
    *q = (*q)[1:]
    return v
}

func uniquePaths(m int, n int) int {
    total := 0
    start := [2]int{0,0}
    q := newQueue()
    q.Enqueue(start)
    for len(*q) != 0 {
        crds := q.Dequeue()
        if crds[0] == m - 1 && crds[1] == n - 1 {
            total++
            continue
        }
        if crds[0] + 1 < m {
            q.Enqueue([2]int{crds[0]+1, crds[1]})
        }
        if crds[1] + 1 < n {
            q.Enqueue([2]int{crds[0], crds[1]+1})
        }

    }
    return total
}
