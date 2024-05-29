/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maxDiameter := 0
    getDiameterOfTree(root, &maxDiameter)
    return maxDiameter
}

func getDiameterOfTree(node *TreeNode, maxDiameter *int) int {
    diameter := 0
    if node.Left != nil {
        diameter += getDepthOfTree(node.Left)
        getDiameterOfTree(node.Left, maxDiameter)
    }
    if node.Right != nil {
        diameter += getDepthOfTree(node.Right)
        getDiameterOfTree(node.Right, maxDiameter)
    }
    fmt.Printf("DIAMETER %d: %d\n", node.Val, diameter)
    if diameter > *maxDiameter {
        *maxDiameter = diameter
    }
    return diameter
}

func getDepthOfTree(node *TreeNode) int {
    var depthLeft, depthRight int
    if node.Left != nil {
        depthLeft = getDepthOfTree(node.Left)
    }
    if node.Right != nil {
        depthRight = getDepthOfTree(node.Right)
    }
    fmt.Printf("%d: %d\n", node.Val, max(depthLeft, depthRight) + 1)
    return max(depthLeft, depthRight) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
