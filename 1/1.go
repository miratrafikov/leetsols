func twoSum(nums []int, target int) []int {
    numToIndex := make(map[int]int, 0)
    if len(nums) == 2 {
        return []int{0, 1}
    }
    var result []int
    for i, num := range(nums) {
        remainder := target - num
        if remainderIndex, ok := numToIndex[remainder]; ok {
            result = []int{remainderIndex, i}
            break
        } else {
            numToIndex[num] = i
        }
    }
    return result
}
