// This time Lomuto's scheme is used
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
    // therefore, if m == targetIndex then all elements up to and including m
    // are guaranteed to be the smallest m elements in the kClosest's input array
    m := partition(points, l, r)
    if m > targetIndex  {
        doQuickselectPartitions(points, l, m-1, targetIndex)
    } else if m < targetIndex {
        doQuickselectPartitions(points, m+1, r, targetIndex)
    }
}

func partition(points [][]int, l, r int) int {
    pivotPoint := points[r]
    i := l - 1
    for j := l; j < r; j++ {
        if less(points[j], pivotPoint) {
            i++
            points[i], points[j] = points[j], points[i]
        }
    }
    i++
    points[i], points[r] = points[r], points[i]
    return i
}

func less(a, b []int) bool {
    distanceA := a[0]*a[0] + a[1]*a[1]
    distanceB := b[0]*b[0] + b[1]*b[1]
    return distanceA < distanceB
}
