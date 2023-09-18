# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        depthToVals = {}
        self.getValsOnAllLevels(root, 0, depthToVals)
        depthToAverageVal = [None]*len(depthToVals)
        for depth, vals in depthToVals.items():
            sumOfVals = 0
            for val in vals:
                sumOfVals += val
            averageVal = sumOfVals / len(vals)
            depthToAverageVal[depth] = averageVal
        return depthToAverageVal
    
    def getValsOnAllLevels(self, root: Optional[TreeNode], depth: int, depthToVals: Dict) -> None:
        if root is None:
            return
        if depth not in depthToVals:
            depthToVals[depth] = [root.val]
        else:
            depthToVals[depth].append(root.val)
        self.getValsOnAllLevels(root.left, depth + 1, depthToVals)
        self.getValsOnAllLevels(root.right, depth + 1, depthToVals)
