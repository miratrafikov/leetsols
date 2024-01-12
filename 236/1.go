/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    routeToP := getRouteToNode(root, p, []*TreeNode{})
    routeToQ := getRouteToNode(root, q, []*TreeNode{})
    i := 0
    for i < len(routeToP) && i < len(routeToQ) && routeToP[i] == routeToQ[i] {
        i++
    }
    return routeToP[i-1]
}

func getRouteToNode(currentNode, targetNode *TreeNode, route []*TreeNode) []*TreeNode {
    if currentNode == nil {
        return route
    }
    route = append(route, currentNode)
    if currentNode == targetNode {
        return route
    }
    routeFromLeft := getRouteToNode(currentNode.Left, targetNode, route)
    if routeFromLeft[len(routeFromLeft)-1] == targetNode {
        return routeFromLeft
    }
    return getRouteToNode(currentNode.Right, targetNode, route)
}

func min(a, b int) int  {
    if a < b {
        return a
    }
    return b
}
