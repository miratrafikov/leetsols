# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findTarget(self, root: Optional[TreeNode], key: int) -> (Optional[TreeNode], Optional[TreeNode]):
        prev = None
        curr = root
        while curr:
            if curr.val == key:
                break
            prev = curr
            if curr.val > key:
                curr = curr.left
            else:
                curr = curr.right
        return curr, prev
    
    def findReplacer(self, root) -> (Optional[TreeNode], Optional[TreeNode]):
        prev = None
        curr = root
        # root has no children
        if not curr.left and not curr.right:
            return curr, prev
        # if right subtree exists, find node with minimum value on the right subtree
        if curr.right:
            prev = curr
            curr = curr.right
            while curr.left:
                prev = curr
                curr = curr.left
        # else find node with maximum value on the left subtree
        else:
            prev = curr
            curr = curr.left
            while curr.right:
                prev = curr
                curr = curr.right
        return curr, prev
    
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        node_delete, node_delete_prev = self.findTarget(root, key)
        # if no node with the specified key found
        if not node_delete:
            return root
        # else find a suitable replacement
        node_replacer, node_replacer_prev = self.findReplacer(node_delete)
        # if deletion node has no children (is a leaf)
        if node_replacer == node_delete:
            # if deletion node is the root, return an empty BST
            if not node_delete_prev:
                return None
            # else delete the reference in the parent of the replacement node
            if node_delete_prev.right == node_delete:
                node_delete_prev.right = None
            else:
                node_delete_prev.left = None
        # else, damn it took me long to find a bug here
        else:
            node_delete.val = node_replacer.val
            if node_replacer_prev.right == node_replacer:
                if node_replacer_prev == node_delete:
                    node_replacer_prev.right = node_replacer.right
                else:
                    node_replacer_prev.right = node_replacer.left
            else:
                if node_replacer_prev == node_delete:
                    node_replacer_prev.left = node_replacer.left
                else:
                    node_replacer_prev.left = node_replacer.right
        return root
