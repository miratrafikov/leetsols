func threeSum(nums []int) [][]int {
    n := len(nums)
    quicksort(nums, 0, n-1)
    triplets := make([][]int, 0)
    for i := 0; i < n - 2; i++ {
        if i >= 1 && nums[i] == nums[i-1] {
            continue
        }
        l := i + 1
        r := n - 1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            switch {
            case sum == 0:
                triplets = append(triplets, []int{nums[i], nums[l], nums[r]})
                l++
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
            case sum < 0:
                l++
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
            case sum > 0:
                r--
                for r > l && nums[r] == nums[r+1] {
                    r--
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
