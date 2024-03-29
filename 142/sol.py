# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        visited = { head }
        
        if not head:
            return None
        
        while head.next:
            head = head.next

            if head in visited:
                return head
            else:
                visited.add(head)
        
        return None
