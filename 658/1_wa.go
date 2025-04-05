func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    dists := make([]int, n)
    for i := 0; i < n; i++ {
        dists[i] = arr[i] - x
        if dists[i] < 0 {
            dists[i] = -dists[i]
        }
    }
    min := 0
    for i := 0; i < n; i++ {
        if dists[i] < dists[min] {
            min = i
        }
    }
    closest := []int{arr[min]}
    l, r := min - 1, min + 1
    for len(closest) != k {
        if l >= 0 && r < n {
            if dists[l] < dists[r] {
                closest = append(closest, arr[l])
                l--
            } else if dists[r] > dists[l] {
                closest = append(closest, arr[r])
                r++
            } else {
                closest = append(closest, arr[l])
                l--
            }
        } else if l >= 0 {
            closest = append(closest, arr[l])
            l--
        } else {
            closest = append(closest, arr[r])
            r++
        }
    }
    quicksort(closest, 0, k-1)
    return closest
}

func quicksort(arr []int, l, r int) {
    if l >= r {
        return
    }
    pivot := partition(arr, l, r)
    quicksort(arr, l, pivot)
    quicksort(arr, pivot+1, r)
}

// Hoare's algorithm
func partition(arr []int, l, r int) int {
    pivot := arr[l + (r - l) / 2]
    i := l - 1
    j := r + 1
    for {
        for {
            i++
            if arr[i] >= pivot {
                break
            }
        }
        for {
            j--
            if arr[j] <= pivot {
                break
            }
        }
        if i >= j {
            break
        }
        arr[i], arr[j] = arr[j], arr[i]
    }
    return j
}
