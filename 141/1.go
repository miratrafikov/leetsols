/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    fast, slow := head, head
    i := 0
    for fast != nil {
        fast = fast.Next
        if i % 2 == 1 {
            slow = slow.Next
        }
        if fast == slow {
            return true
        }   
        i++
    }
    return false
}
