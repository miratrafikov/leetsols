# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque

class Solution:
    def isSameTree(self, root_a: Optional[TreeNode], root_b: Optional[TreeNode]) -> bool:
        queue_a, queue_b = deque([root_a]), deque([root_b])

        while queue_a and queue_b:
            node_a, node_b = queue_a.popleft(), queue_b.popleft()

            if not node_a and not node_b:
                continue
            
            if not node_a or not node_b:
                return False

            if node_a.val != node_b.val:
                return False
            
            queue_a.append(node_a.left)
            queue_a.append(node_a.right)
            queue_b.append(node_b.left)
            queue_b.append(node_b.right)
        
        return True
