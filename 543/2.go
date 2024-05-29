/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    getDepthOfTreeAndUpdateMaxDiameter(root, &maxDiameter)
    return maxDiameter
}

func getDepthOfTreeAndUpdateMaxDiameter(node *TreeNode, maxDiameter *int) int {
    var depthLeft, depthRight int
    diameter := 0
    if node.Left != nil {
        depthLeft = getDepthOfTreeAndUpdateMaxDiameter(node.Left, maxDiameter)
        diameter += depthLeft
    }
    if node.Right != nil {
        depthRight = getDepthOfTreeAndUpdateMaxDiameter(node.Right, maxDiameter)
        diameter += depthRight
    }
    *maxDiameter = max(*maxDiameter, diameter)
    return max(depthLeft, depthRight) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
