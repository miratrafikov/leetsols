class Solution:
    def maxDistance(self, colors: List[int]) -> int:
        l = [None] * 101
        for i in range(len(colors)):
            if l[colors[i]]:
                l[colors[i]][1] = i
            else:
                l[colors[i]] = [i, i]
        ans = 0
        for i in range(len(l) - 1):
            if not l[i]:
                continue
            for j in range(i + 1, len(l)):
                if not l[j]:
                    continue
                diff_l = abs(l[i][1] - l[j][0])
                if ans < diff_l:
                    ans = diff_l
                diff_r = abs(l[i][0] - l[j][1])
                if ans < diff_r:
                    ans = diff_r
	        return ans
