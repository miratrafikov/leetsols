/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    dedup := head
    for dedup != nil  {
        runner := dedup.Next
        for runner != nil && runner.Val == dedup.Val {
            runner = runner.Next
        }
        dedup.Next = runner
        dedup = runner
    }
    return head
}
