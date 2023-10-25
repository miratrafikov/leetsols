/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head.Next == nil {
        return head
    }
    slow, fast := head, head
    shouldSlowMove := false
    for fast != nil {
        fast = fast.Next
        if shouldSlowMove {
            slow = slow.Next
        }
        shouldSlowMove = !shouldSlowMove
    }
    return slow
}
