func kClosest(points [][]int, k int) [][]int {
    // if length of the array in N and we need to return K <= N smallest elements
    // (where a is smaller than b if its Euclidean distance to origin is smaller)
    // then index up to which are the elements of interest will be stored in a partially sorted array
    // that is made by quickselect partitions is K-1
    targetIndex := k - 1
    doQuickselectPartitions(points, 0, len(points)-1, targetIndex)
    return points[:targetIndex+1]
}

func doQuickselectPartitions(points [][]int, l, r, targetIndex int) {
    if l == r {
        return
    }
    // elements to the left of the points[m] are less or equal to it
    // elements to the right of the points[m] are greater or equal to it
    m := partition(points, l, r)
    // (note for myself) reason why we don't return if m == targetIndex is
    // because the left part is not sorted, even though all elements to the left are lesser or equal
    if m >= targetIndex  {
        doQuickselectPartitions(points, l, m, targetIndex)
    } else {
        doQuickselectPartitions(points, m+1, r, targetIndex)
    }
}

func partition(points [][]int, l, r int) int {
    pivotPoint := points[l + (r - l) / 2]
    i := l - 1
    j := r + 1
    for {
        for {
            i++
            if !less(points[i], pivotPoint) {
                break
            }
        }
        for {
            j--
            if !less(pivotPoint, points[j]) {
                break
            }
        }
        if i >= j {
            break
        }
        points[i], points[j] = points[j], points[i]
    }
    return j
}

func less(a, b []int) bool {
    distanceA := math.Sqrt(math.Pow(float64(a[0]), 2) + math.Pow(float64(a[1]), 2))
    distanceB := math.Sqrt(math.Pow(float64(b[0]), 2) + math.Pow(float64(b[1]), 2))
    return distanceA < distanceB
}
