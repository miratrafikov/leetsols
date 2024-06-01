func sortColors(nums []int)  {
    n := len(nums)
    m := map[int]int{
        0: 0,
        1: 0,
        2: 0,
    }
    for i := 0; i < n; i++ {
        m[nums[i]]++
    }
    i := 0
    for range m[0]  {
        nums[i] = 0
        i++
    }
    for range m[1] {
        nums[i] = 1
        i++
    }
    for range m[2] {
        nums[i] = 2
        i++
    }
}
