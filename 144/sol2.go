/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    result := []int{}
    if root == nil {
        return result
    }
    stack := make(tnStack, 0)
    stack = stack.Push(root)
    var node *TreeNode
    for len(stack) > 0 {
        stack, node = stack.Pop()
        result = append(result, (node).Val)
        if node.Right != nil {
            stack = stack.Push(node.Right)
        }
        if node.Left != nil {
            stack = stack.Push(node.Left)
        }
    }
    return result
}

type tnStack []*TreeNode

func (s tnStack) Push(x *TreeNode) tnStack {
    return append(s, x)
}

func (s tnStack) Pop() (tnStack, *TreeNode) {
    return s[:len(s) - 1], s[len(s) - 1]
}
