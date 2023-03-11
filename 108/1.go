/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return aboba(nums, 0, len(nums) - 1)
}

func aboba(nums[] int, p, r int) *TreeNode {
    if p > r {
        return nil
    }

    q := p + (r - p) / 2
    node := TreeNode{Val: nums[q]}
    node.Left = aboba(nums, p, q - 1)
    node.Right = aboba(nums, q + 1, r)

    return &node
}
