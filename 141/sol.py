# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if not head:
            return False
        fast = head
        slow = head
        slow_move = True
        while fast.next:
            fast = fast.next
            if fast == slow:
                return True
            slow_move = not slow_move
            if slow_move:
                slow = slow.next
        return False 
