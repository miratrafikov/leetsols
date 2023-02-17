/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    values := []int{}
    collectValues(root, &values)

    minDiff := values[1] - values[0]
    for i := 2; i < len(values); i++ {
        currDiff := values[i] - values[i - 1]
        if currDiff < minDiff {
            minDiff = currDiff
        }
    }

    return minDiff
}

func collectValues(root *TreeNode, values *[]int) {
    if root == nil {
        return
    }

    collectValues(root.Left, values)
    *values = append(*values, root.Val)
    collectValues(root.Right, values)
}
