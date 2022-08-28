# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l: Optional[ListNode], r: Optional[ListNode]) -> Optional[ListNode]:
        if not l:
            return r
        
        if not r:
            return l
            
        result_head = None
        result_tail = None
        
        if l.val < r.val:
            result_head = result_tail = l
            l = l.next
        else:
            result_head = result_tail = r
            r = r.next
        
        while l and r:
            if l.val < r.val:
                result_tail.next = l
                result_tail = result_tail.next
                l = l.next
            else:
                result_tail.next = r
                result_tail = result_tail.next
                r = r.next
        
        while l:
            result_tail.next = l
            result_tail = result_tail.next
            l = l.next
            
        while r:
            result_tail.next = r
            result_tail = result_tail.next
            r = r.next
        
        return result_head
