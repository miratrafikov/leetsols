class Solution:
    def longestPalindrome(self, s: str) -> int:
        asciiToOccurences = [0]*(ord('z')-ord('A')+1)
        for c in s:
            asciiToOccurences[ord(c)-ord('A')] += 1
        palindromeLength = 0
        for letterOccurences in asciiToOccurences:
            palindromeLength += letterOccurences // 2 * 2
            if palindromeLength % 2 == 0:
                palindromeLength += letterOccurences % 2
        return palindromeLength
