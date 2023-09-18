# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        queue_1 = [root]
        queue_2 = []
        queue_pops_count = 0
        level_to_average_val = []
        curr_level = 0
        curr_level_sum = 0
        while len(queue_1) > 0:
            node = queue_1.pop(0)
            queue_pops_count += 1
            curr_level_sum += node.val
            if node.left:
                queue_2.append(node.left)
            if node.right:
                queue_2.append(node.right)
            if len(queue_1) == 0:
                level_to_average_val.append(curr_level_sum / queue_pops_count)
                curr_level += 1
                curr_level_sum = 0
                queue_pops_count = 0
                queue_1 = queue_2
                queue_2 = []
        return level_to_average_val
