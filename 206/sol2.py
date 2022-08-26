# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, curr: Optional[ListNode]) -> Optional[ListNode]:
        if not curr or not curr.next:
            return curr
        
        next = self.reverseList(curr.next)
        curr.next.next = curr
        curr.next = None

        return next
