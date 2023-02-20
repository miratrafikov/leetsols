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
    var levelOutput []int
    var queue1 []*TreeNode
    var queue2 []*TreeNode

    enqueue(&queue1, root)

    for len(queue1) > 0 || len(queue2) > 0 {
        for len(queue1) > 0 {
            node := deque(&queue1)

            if node != nil {
                levelOutput = append(levelOutput, node.Val)
                queue2 = append(queue2, node.Left)
                queue2 = append(queue2, node.Right)
            }
        }

        if len(levelOutput) > 0 {
            totalOutput = append(totalOutput, levelOutput)
            levelOutput = make([]int, 0)
        }

        for len(queue2) > 0 {
            node := deque(&queue2)

            if node != nil {
                levelOutput = append(levelOutput, node.Val)
                queue1 = append(queue1, node.Left)
                queue1 = append(queue1, node.Right)
            }
        }

        if len(levelOutput) > 0 {
            totalOutput = append(totalOutput, levelOutput)
            levelOutput = make([]int, 0)
        }
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
