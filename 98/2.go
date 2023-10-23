/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return soIsIt(root.Left, math.MinInt, root.Val) && soIsIt(root.Right, root.Val, math.MaxInt)
}

func soIsIt(root *TreeNode, leftBoundary, rightBoundary int) bool {
    if root == nil {
        return true
    }
    if root.Val >= rightBoundary || root.Val <= leftBoundary {
        return false
    }
    return soIsIt(root.Left, leftBoundary, root.Val) && soIsIt(root.Right, root.Val, rightBoundary)
}
