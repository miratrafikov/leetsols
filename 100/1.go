/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dequeue(q *[]*TreeNode) *TreeNode {
    v := (*q)[0]
    *q = (*q)[1:]

    return v
}

func isSameTree(rootA *TreeNode, rootB *TreeNode) bool {
    var queueA []*TreeNode
    var queueB []*TreeNode
    queueA = append(queueA, rootA)
    queueB = append(queueB, rootB)

    for len(queueA) > 0 {
        nodeA := dequeue(&queueA)
        nodeB := dequeue(&queueB)

        if nodeA == nil && nodeB == nil {
            continue
        }

        if nodeA == nil || nodeB == nil {
            return false
        }

        if (*nodeA).Val != (*nodeB).Val {
            return false
        }

        queueA = append(queueA, (*nodeA).Left)
        queueA = append(queueA, (*nodeA).Right)
        queueB = append(queueB, (*nodeB).Left)
        queueB = append(queueB, (*nodeB).Right)
    }
    
    return true
}
