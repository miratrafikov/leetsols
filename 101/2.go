/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return juxtaposeBranches(root.Left, root.Right)
}

func juxtaposeBranches(l, r *TreeNode) bool {
    if l == nil && r == nil {
        return true
    }
    if l == nil || r == nil || l.Val != r.Val {
        return false
    }
    return juxtaposeBranches(l.Left, r.Right) && juxtaposeBranches(l.Right, r.Left)
}
