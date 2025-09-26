func triangleNumber(nums []int) int {
    n := len(nums)

    quicksort(nums, 0, n-1)

    // skip all zeroes at the beginning
    trueLeft := 0
    for trueLeft < n && nums[trueLeft] == 0 {
        trueLeft++
    }

    count := 0
    for k := n-1; k >= trueLeft + 2; k-- {
        i := trueLeft
        j := k - 1
        for i < j {
            if nums[i] + nums[j] > nums[k] {
                count += j - i
                j--
            } else {
                i++
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
