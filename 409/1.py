class Solution:
    def longestPalindrome(self, s: str) -> int:
        m = {}
        for c in s:
            if c in m:
                m[c] += 1
            else:
                m[c] = 1
        palindromeLength = 0
        orphanLetterMet = False
        for k, v in m.items():
            if v % 2 == 0:
                palindromeLength += v
            else:
                palindromeLength += v - 1
                orphanLetterMet = True
        if orphanLetterMet:
            return palindromeLength + 1
        return palindromeLength
