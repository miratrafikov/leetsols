/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    // return nil if the list is empty
    if head == nil {
        return nil
    }
    if k == 0 {
        return head
    }
    // determine length of the list and the minimal effective number of repositions
    n := getLengthOfList(head)
    rotationsNeeded := k % n
    if rotationsNeeded == 0 {
        return head
    }
    // now that we know N and "true" k, we now need to:
    // 1. make element at N-K-1 the end of the list
    // 2. connect the previous end of the list to head
    // 3. make element at N-K the new head of the list
    currentElementIndex := 0
    var prev *ListNode
    tail := head
    for currentElementIndex != n - rotationsNeeded {
        prev = tail
        tail = tail.Next
        currentElementIndex++
    }
    prev.Next = nil
    newHead := tail
    for tail.Next != nil {
        tail = tail.Next
    }
    tail.Next = head
    return newHead
}

func getLengthOfList(head *ListNode) int {
    length := 1
    tail := head
    for tail.Next != nil {
        tail = tail.Next
        length++
    }
    return length
}
