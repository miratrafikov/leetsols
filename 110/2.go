/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return determineMaxDepth(root) != -1
}

func determineMaxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    depthLeft := determineMaxDepth(root.Left)
    depthRight := determineMaxDepth(root.Right)

    if depthLeft == -1 || depthRight == -1 || abs(depthLeft - depthRight) > 1 {
        return -1
    }

    return 1 + max(depthLeft, depthRight)
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
