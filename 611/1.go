func triangleNumber(nums []int) int {
    n := len(nums)
    quicksort(nums, 0, n-1)
    count := 0
    for i := 0; i < n - 2; i++ {
        if nums[i] == 0 {
            continue
        }
        for j := i + 1; j < n - 1; j++ {
            maxThirdSide := nums[i] + nums[j] - 1
            for k := n-1; k > j; k-- {
                if nums[k] <= maxThirdSide {
                    count += k - j
                    break
                }
            }
        }
    }
    return count
}

func quicksort(arr []int, l, r int) {
    if l >= r {
        return
    }
    p := partitionHoare(arr, l, r)
    quicksort(arr, l, p)
    quicksort(arr, p+1, r)
}

func partitionHoare(arr []int, l, r int) int {
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
