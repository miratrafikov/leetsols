/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type SumAndFrequency struct {
    Sum int
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
    maxFrequency := -1
    var numbersWithMaxFrequency []int
    for num, frequency := range numToFrequency {
        if frequency > maxFrequency {
            maxFrequency = frequency
            numbersWithMaxFrequency = []int{num}
        } else if frequency == maxFrequency {
            numbersWithMaxFrequency = append(numbersWithMaxFrequency, num)
        }
    }
    return numbersWithMaxFrequency
}
