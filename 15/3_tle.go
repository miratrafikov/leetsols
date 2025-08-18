func threeSum(nums []int) [][]int {
    n := len(nums)
    quicksort(nums, 0, n-1)
    // m stored occurence count for every number
    m := make(map[int]int)
    // nsu = nums unique sorted
    nsu := []int{nums[0]}
    for i := 0; i < n; i++ {
        m[nums[i]]++
        if nsu[len(nsu)-1] != nums[i] {
            nsu = append(nsu, nums[i])
        }
    }
    triplets := make([][]int, 0)
    for i := 0; i < len(nsu); i++ {
        x := nsu[i]
        jStart := i + 1
        if m[x] >= 2 {
            jStart = i
        }
        for j := jStart; j < len(nsu); j++ {
            y := nsu[j]
            kStart := j + 1
            if m[y] >= 2 && y != x {
                kStart = j
            }
            if m[y] >= 3 && y == x {
                kStart = j
            }
            for k := kStart; k < len(nsu); k++ {
                z := nsu[k]
                if x + y + z == 0 {
                    isTripletKnown := false
                    if len(triplets) > 0 {
                        lastKnownTriplet := triplets[len(triplets)-1]
                        isTripletKnown = lastKnownTriplet[0] == x && 
                            lastKnownTriplet[1] == y &&
                            lastKnownTriplet[2] == z
                    }
                    if !isTripletKnown {
                        triplets = append(triplets, []int{x,y,z})
                    }
                }
            }
        }
    }
    return triplets
}

func quicksort(arr []int, l, r int) {
    if l >= r {
        return
    }
    p := hoarePartition(arr, l, r)
    quicksort(arr, l, p)
    quicksort(arr, p+1, r)
}

func hoarePartition(arr []int, l, r int)  int {
    pivot := arr[l + (r - l) / 2]
    i := l-1
    j := r+1
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
