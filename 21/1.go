/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // create a pointer to the head of the resulting list (the head is what the function will return)
    var head *ListNode
    // create a pointer to the tail of the resuling list
    var tail *ListNode
    // initialize head
    if list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            head = list1
            list1 = list1.Next
        } else {
            head = list2
            list2 = list2.Next
        }
    } else if list1 != nil {
        head = list1
        list1 = list1.Next
    } else if list2 != nil {
        head = list2
        list2 = list2.Next
    } else {
        return head
    }
    // resulting list has length of 1 so tail should point to head
    tail = head
    // start inserting nodes from the provided lists one by one into the resulting list
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            tail = tail.Next
            list1 = list1.Next
        } else {
            tail.Next = list2
            tail = tail.Next
            list2 = list2.Next
        }
    }
    for list1 != nil {
        tail.Next = list1
        tail = tail.Next
        list1 = list1.Next
    }
    for list2 != nil {
        tail.Next = list2
        tail = tail.Next
        list2 = list2.Next
    }
    // return the head of the resulting list
    return head
}
