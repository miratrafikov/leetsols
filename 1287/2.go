func findSpecialInteger(arr []int) int {
    n := len(arr)
    targetOccurences := n / 4 + 1
    positionsOfCandidates := []int{n/4, n/2, 3*n/4}
    var specialElement int
    for _, candidatePosition := range positionsOfCandidates {
        leftmostPosition := findIndexOfLeftmostOccurence(arr, candidatePosition)
        rightmostPosition := findIndexOfRightmostOccurence(arr, candidatePosition)
        fmt.Println(leftmostPosition, rightmostPosition)
        if rightmostPosition + 1 - leftmostPosition >= targetOccurences {
            specialElement = arr[candidatePosition]
            break
        }
    }
    return specialElement
}

func findIndexOfLeftmostOccurence(arr []int, r int) int {
    l := 0
    for l < r {
        m := l + (r - l) / 2
        if arr[m] == arr[r] {
            r = m
        } else {
            l = m + 1
        }
    }
    return r
}

func findIndexOfRightmostOccurence(arr []int, l int) int {
    r := len(arr) - 1
    for l < r {
        m := l + (r - l + 1) / 2
        if arr[m] == arr[l] {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}
