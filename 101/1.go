/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    valuesLeft := []int{}
    valuesRight := []int{}
    collectValuesLeft(root.Left, &valuesLeft)
    collectValuesRight(root.Right, &valuesRight)
    return areEqual(valuesLeft, valuesRight)
}

func collectValuesLeft(node *TreeNode, acc *[]int) {
    if node == nil {
        *acc = append(*acc, -101)
        return
    }
    *acc = append(*acc, node.Val)
    collectValuesLeft(node.Left, acc)
    collectValuesLeft(node.Right, acc)
}

func collectValuesRight(node *TreeNode, acc *[]int) {
    if node == nil {
        *acc = append(*acc, -101)
        return
    }
    *acc = append(*acc, node.Val)
    collectValuesRight(node.Right, acc)
    collectValuesRight(node.Left, acc)
}

func areEqual(s1, s2 []int) bool {
    if len(s1) != len(s2) {
        return false
    }
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            return false
        }
    }
    return true
}
