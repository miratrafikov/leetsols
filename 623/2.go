/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    return addRow(root, val, depth, true)
}

func addRow(node *TreeNode, val, depth int, isLeft bool) *TreeNode {
    if depth == 1 {
        if isLeft {
            return &TreeNode{
                Val: val,
                Left: node,
            }
        } else {
            return &TreeNode{
                Val: val,
                Right: node,
            }
        }
    }
    if node != nil {
        node.Left = addRow(node.Left, val, depth-1, true)
        node.Right = addRow(node.Right, val, depth-1, false)
    }
    return node
}
