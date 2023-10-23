/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    isLeftValid := soIsIt(root.Left, -1, root.Val)
    isRightValid := soIsIt(root.Right, root.Val, -1)
    return isLeftValid && isRightValid
}

func soIsIt(root *TreeNode, minValue int, maxValue int) bool {
    if root == nil {
        return true
    }
    if minValue != -1 && root.Val <= minValue {
        return false
    }
    if maxValue != -1 && root.Val >= maxValue {
        return false
    }
    isLeftValid := soIsIt(root.Left, minValue, root.Val)
    isRightValid := soIsIt(root.Right, root.Val, maxValue)
    return isLeftValid && isRightValid
}
