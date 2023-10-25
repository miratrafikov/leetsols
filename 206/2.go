/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    headOfReversedRestOfTheList := reverseList(head.Next)
    tailOfReversedRestOfTheList := head.Next
    tailOfReversedRestOfTheList.Next = head
    head.Next = nil
    return headOfReversedRestOfTheList
}
