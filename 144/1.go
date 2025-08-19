/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type stack []TreeNode

func (s *stack) Push(node TreeNode) {
    *s = append(*s, node)
}

func (s *stack) Pop() TreeNode {
    node := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return node
}

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    s := new(stack)
    s.Push(*root)
    vals := []int{}
    for len(*s) != 0 {
        top := s.Pop()
        vals = append(vals, top.Val)
        if top.Right != nil {
            s.Push(*top.Right)
        }
        if top.Left != nil {
            s.Push(*top.Left)
        }
    }
    return vals
}
