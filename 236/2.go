/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root == q || root == p {
        return root
    }
    searchResultFromLeft := lowestCommonAncestor(root.Left, q, p)
    searchResultFromRight := lowestCommonAncestor(root.Right, q, p)
    if searchResultFromLeft != nil && searchResultFromRight != nil {
        return root
    }
    if searchResultFromLeft != nil {
        return searchResultFromLeft
    }
    if searchResultFromRight != nil {
        return searchResultFromRight
    }
    return nil
}
