/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    depthLeft := determineMaxDepth(root.Left)
    depthRight := determineMaxDepth(root.Right)

    if abs(depthLeft - depthRight) > 1 {
        return false
    }

    return isBalanced(root.Left) && isBalanced(root.Right) 
}

func determineMaxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + max(determineMaxDepth(root.Left), determineMaxDepth(root.Right))
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}
