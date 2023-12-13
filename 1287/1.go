func findSpecialInteger(arr []int) int {
    n := len(arr)
    targetOccurences := n / 4 + 1
    var element int
    for i := 0; i <= n - targetOccurences; i++ {
        if arr[i] == arr[i + targetOccurences - 1] {
            element = arr[i]
            break
        }
    }
    return element
}
