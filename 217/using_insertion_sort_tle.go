func containsDuplicate(nums []int) bool {
    sort(&nums)
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] {
            return true
        }
    }

    return false
}

func sort(nums *[]int) {
    for i := 1; i < len(*nums); i++ {
        original := (*nums)[i]
        j := i - 1
        for ; j >= 0 && (*nums)[j] > original; j-- {
            (*nums)[j + 1]  = (*nums)[j]
        }

        (*nums)[j + 1] = original
    }
}
