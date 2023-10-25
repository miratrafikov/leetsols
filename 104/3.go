/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maxDepthLeft := maxDepth(root.Left)
    maxDepthRight := maxDepth(root.Right)
    return max(maxDepthLeft, maxDepthRight) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
