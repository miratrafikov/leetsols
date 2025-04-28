func combinationSum(candidates []int, target int) [][]int {
    slices.Sort(candidates)
    combinations := [][]int{}
    generate(candidates, 0, target, []int{}, 0, &combinations)
    return combinations
}

func generate(nums []int, i int, target int, acc []int, accSum int, doneAcc *[][]int) {
    if accSum > target {
        return
    }
    if i >= len(nums) {
        return
    }
    generate(nums, i+1, target, acc, accSum, doneAcc)
    accSum += nums[i]
    if accSum > target {
        return
    }
    acc = append(append([]int{}, acc...), nums[i])
    if accSum == target {
        *doneAcc = append(*doneAcc, acc)
        return
    }
    generate(nums, i, target, acc, accSum, doneAcc)
}
