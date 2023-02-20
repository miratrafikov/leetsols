/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var totalOutput [][]int

    if root == nil {
        return totalOutput
    }

    var levelOutput []int
    var queue []*TreeNode

    enqueue(&queue, root)

    for len(queue) > 0 {
        for nodesAtCurrentLevel := len(queue); nodesAtCurrentLevel != 0; nodesAtCurrentLevel-- {
            node := deque(&queue)
            levelOutput = append(levelOutput, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil{
                queue = append(queue, node.Right)
            }
        }

        totalOutput = append(totalOutput, levelOutput)
        levelOutput = make([]int, 0)
    }

    return totalOutput
}

func enqueue(q *[]*TreeNode, node *TreeNode) {
    *q = append(*q, node)
}

func deque(q *[]*TreeNode) *TreeNode {
    node := (*q)[0]
    *q = (*q)[1:]

    return node
}
