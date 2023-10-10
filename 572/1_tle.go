/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    return bruh(root, subRoot, subRoot)
}

func bruh(root *TreeNode, subRoot *TreeNode, trueSubRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
        return true
    }
    if root == nil && subRoot != nil || root != nil && subRoot == nil {
        return false
    }
    if root.Val == subRoot.Val && bruh(root.Left, subRoot.Left, trueSubRoot) && bruh(root.Right, subRoot.Right, trueSubRoot) {
        return true
    }
    return bruh(root.Left, trueSubRoot, trueSubRoot) || bruh(root.Right, trueSubRoot, trueSubRoot)
}
