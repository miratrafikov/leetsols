/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeProperties struct {
    IsBalanced bool
    Height     int
}

 
func isBalanced(root *TreeNode) bool {
    return utility(root).IsBalanced
}

func utility(root *TreeNode) TreeProperties {
    if root == nil {
        return TreeProperties{true, 0}
    }

    leftTreeProperties := utility(root.Left)
    rightTreeProperties := utility(root.Right)

    isBalanced := leftTreeProperties.IsBalanced &&
        rightTreeProperties.IsBalanced &&
        abs(leftTreeProperties.Height - rightTreeProperties.Height) <= 1

    return TreeProperties{
        isBalanced,
        1 + max(leftTreeProperties.Height, rightTreeProperties.Height),
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
