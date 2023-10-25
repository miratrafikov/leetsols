/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    pathToP := pathToTarget(root, p, []*TreeNode{root})
    pathToQ := pathToTarget(root, q, []*TreeNode{root})
    i := 0
    for i < len(pathToP) && i < len(pathToQ) {
        if pathToP[i] != pathToQ[i] {
            break
        }
        i++
    }
    return pathToP[i-1]
}

func pathToTarget(root, targetNode *TreeNode, path []*TreeNode) []*TreeNode {
    path = append(path, root)
    if root == targetNode {
        return path
    }
    if root == nil {
        return []*TreeNode{}
    }
    pathLeft := pathToTarget(root.Left, targetNode, path)
    pathRight := pathToTarget(root.Right, targetNode, path)
    if len(pathLeft) != 0 {
        return pathLeft
    }
    if len(pathRight) != 0 {
        return pathRight
    }
    return []*TreeNode{}
}
