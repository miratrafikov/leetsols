/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type NumberAndFrequency struct {
    Number int
    Frequency int
}

func findFrequentTreeSum(root *TreeNode) []int {
    treeSums := make([]int, 0)
    addTreeSumToTreeSumsAndReturnIt(root, &treeSums)
    numberToFrequency := mapNumbersToFrequencies(&treeSums)
    return getNumbersWithTopFrequency(numberToFrequency)
}

func addTreeSumToTreeSumsAndReturnIt(root *TreeNode, treeSums *[]int) int {
    leftSubtreeSum := 0
    rightSubtreeSum :=  0
    if root.Left != nil {
        leftSubtreeSum = addTreeSumToTreeSumsAndReturnIt(root.Left, treeSums)
    }
    if root.Right != nil  {
        rightSubtreeSum = addTreeSumToTreeSumsAndReturnIt(root.Right, treeSums)
    }
    treeSum := root.Val + leftSubtreeSum +  rightSubtreeSum
    *treeSums = append(*treeSums, treeSum)
    return treeSum
}

func mapNumbersToFrequencies(numbers *[]int) map[int]int {
    m := make(map[int]int)
    for _, num := range *numbers {
        if frequency, isNumInMap := m[num]; isNumInMap {
            m[num] = frequency + 1
        } else {
            m[num] = 1
        }
    }
    return m
}

func getNumbersWithTopFrequency(numToFrequency map[int]int) []int {
    numbersAndFrequencies := make([]NumberAndFrequency, len(numToFrequency))
    i := 0
    for num, frequency := range numToFrequency {
        numbersAndFrequencies[i] = NumberAndFrequency{Number: num, Frequency: frequency}
        i++
    }
    sortNumbersByFrequency(&numbersAndFrequencies)
    topFrequency := numbersAndFrequencies[0].Frequency
    numbersWithTopFrequency := []int{numbersAndFrequencies[0].Number}
    for i := 1; i < len(numbersAndFrequencies) && numbersAndFrequencies[i].Frequency == topFrequency; i++ {
        numbersWithTopFrequency = append(numbersWithTopFrequency, numbersAndFrequencies[i].Number)
    }
    return numbersWithTopFrequency
}

func sortNumbersByFrequency(s *[]NumberAndFrequency) {
    mergesort(s, 0, len(*s)-1)
}

func mergesort(s *[]NumberAndFrequency, l, r int) {
    if r - l == 0 {
        return
    }
    m := l + (r - l) / 2
    mergesort(s, l, m)
    mergesort(s, m + 1, r)
    merge(s, l, m, r)
}

func merge(s *[]NumberAndFrequency, l, m, r int) {
    sorted := make([]NumberAndFrequency, r - l + 1)
    i := 0
    j := 0
    for k := 0; k < len(sorted); k++ {
        if l + i <= m && m + 1 + j <= r {
            leftSubarrayElement := (*s)[l + i]
            rightSubarrayElement := (*s)[m + 1 + j]
            if leftSubarrayElement.Frequency > rightSubarrayElement.Frequency {
                sorted[k] = leftSubarrayElement
                i++
            } else {
                sorted[k] = rightSubarrayElement
                j++
            }
        } else if l + i <= m {
            leftSubarrayElement := (*s)[l + i]
            sorted[k] = leftSubarrayElement
            i++
        } else {
            rightSubarrayElement := (*s)[m + 1 + j]
            sorted[k] = rightSubarrayElement
            j++
        }
    }
    for k := 0; k < len(sorted); k++ {
        (*s)[l + k] = sorted[k]
    }
}
