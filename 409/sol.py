class Solution:
    def longestPalindrome(self, s: str) -> int:
        unpaired_chars = set()
        
        for c in s:
            if c not in unpaired_chars:
                unpaired_chars.add(c)
            else:
                unpaired_chars.remove(c)
        
        usable_letters = len(s) - len(unpaired_chars)
        
        return usable_letters + (0 if len(unpaired_chars) == 0 else 1)
