/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    tail := result
    carry := 0
    for l1 != nil && l2 != nil {
        sum := l1.Val + l2.Val + carry
        tail.Val = sum % 10
        carry = sum / 10
        l1 = l1.Next
        l2 = l2.Next
        if l1 != nil || l2 != nil {
            tail.Next = &ListNode{}
            tail = tail.Next
        }
    }
    for l1 != nil {
        sum := l1.Val + carry
        tail.Val = sum % 10
        carry = sum / 10
        l1 = l1.Next
        if l1 != nil {
            tail.Next = &ListNode{}
            tail = tail.Next
        }
    }
    for l2 != nil {
        sum := l2.Val + carry
        tail.Val = sum % 10
        carry = sum / 10
        l2 = l2.Next
        if l2 != nil {
            tail.Next = &ListNode{}
            tail = tail.Next
        }
    }
    if carry > 0 {
        tail.Next = &ListNode{}
        tail = tail.Next
        tail.Val = carry
    }
    return result
}
