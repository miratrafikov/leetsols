/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    valuesVisibleFromTheRight := []int{}
    if root == nil {
        return valuesVisibleFromTheRight
    }
    currentLevel := []*TreeNode{root}
    nextLevel := []*TreeNode{}
    i := 0
    for i < len(currentLevel) {
        node := currentLevel[i]
        if node.Left != nil {
            nextLevel = append(nextLevel, node.Left)
        }
        if node.Right != nil {
            nextLevel = append(nextLevel, node.Right)
        }
        if i == len(currentLevel) - 1 {
            valuesVisibleFromTheRight = append(valuesVisibleFromTheRight, node.Val)
            currentLevel = nextLevel
            nextLevel = []*TreeNode{}
            i = 0
        } else {
            i++
        }
    }
    return valuesVisibleFromTheRight
}

