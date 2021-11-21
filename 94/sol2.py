from collections import deque


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        result = []
        stack = deque()
        node = root
        while True:
            if node:
                stack.append(node)
                node = node.left
            elif len(stack):
                node = stack.pop()
                result.append(node.val)
                node = node.right
            else:
                break
        return result
