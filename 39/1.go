func combinationSum(candidates []int, target int) [][]int {
    slices.Sort(candidates)
    combinations := make([][]int, 0, 150)
    generate(candidates, 0, target, []int{}, &combinations)
    return combinations
}

func generate(nums []int, i int, target int, acc []int, doneAcc *[][]int) {
    if sum(acc) > target {
        return
    }
    if i >= len(nums) {
        return
    }
    generate(nums, i+1, target, acc, doneAcc)
    acc = append(append([]int{}, acc...), nums[i])
    if sum(acc) == target {
        *doneAcc = append(*doneAcc, acc)
        return
    }
    generate(nums, i, target, acc, doneAcc)
}

func sum(arr []int) int {
    result := 0
    for i := 0; i < len(arr); i++ {
        result += arr[i]
    }
    return result
}
