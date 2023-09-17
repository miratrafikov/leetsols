/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        return &TreeNode{
            Val: val,
            Left: root,
        }
    }
    root.Left = addRow(root.Left, val, depth-1, true)
    root.Right = addRow(root.Right, val, depth-1, false)
    return root
}

func addRow(node *TreeNode, val, depth int, isLeft bool) *TreeNode {
    if depth == 1 {
        newNode := &TreeNode{
            Val: val,
        }
        if isLeft {
            newNode.Left = node
        } else {
            newNode.Right = node
        }
        return newNode
    }
    if node != nil {
        node.Left = addRow(node.Left, val, depth-1, true)
        node.Right = addRow(node.Right, val, depth-1, false)
    }
    return node
}
