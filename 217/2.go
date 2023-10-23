func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, num := range nums {
        if _, seen := m[num]; seen {
            return true
        }
        m[num] = true
    }
    return false
}
