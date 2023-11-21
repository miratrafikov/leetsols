// using sorting first
// time: O(nlogn)
// space: O(logn) as there are logn + X recursive calls each using O(1) space
func kClosest(points [][]int, k int) [][]int {
    quicksort(points, 0, len(points)-1)
    targetIndex := k - 1
    return points[:targetIndex+1]
}

func quicksort(points [][]int, l, r int) {
    if l == r {
        return
    }
    m := partition(points, l, r)
    quicksort(points, l, m)
    quicksort(points, m+1, r)
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
