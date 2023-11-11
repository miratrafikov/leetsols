# Solution 3: shrinking-expanding window with set
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_substring_length = 0
        current_substring_char_set = set()
        i, j = 0, 0
        while j < len(s):
            while s[j] in current_substring_char_set:
                current_substring_char_set.remove(s[i])
                i += 1
            current_substring_char_set.add(s[j])
            if j - i + 1 > max_substring_length:
                max_substring_length = j - i + 1
            j += 1
        return max_substring_length 
