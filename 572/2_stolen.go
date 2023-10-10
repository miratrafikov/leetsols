/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
        return true
    }
    if root == nil && subRoot != nil || root != nil && subRoot == nil {
        return false
    }
    return isEqual(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot);
}

func isEqual(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
        return true
    }
    if root == nil && subRoot != nil || root != nil && subRoot == nil {
        return false
    }
    return root.Val == subRoot.Val && isEqual(root.Left, subRoot.Left) && isEqual(root.Right, subRoot.Right)
}
