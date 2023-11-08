class Solution:
    def findChampion(self, grid: List[List[int]]) -> int:
        n = len(grid)
        for teamToTest in range(n):
            didEverLose = False
            for rivalTeam in range(n):
                if grid[rivalTeam][teamToTest] == 1:
                    didEverLose = True
                    break
            if not didEverLose:
                return teamToTest
