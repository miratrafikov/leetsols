class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if len(s) == 0:
            return True
        i_s = 0
        for i_t in range(len(t)):
            if t[i_t] == s[i_s]:
                i_s += 1
                if i_s == len(s):
                    return True
        return False
