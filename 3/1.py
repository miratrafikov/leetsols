class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_substring_length = 0
        current_substring_length = 0
        current_substring_chars = set()
        i = 0
        while i < len(s):
            if s[i] in current_substring_chars:
                i -= current_substring_length
                current_substring_length = 0
                current_substring_chars = set()
            else:
                current_substring_length += 1
                current_substring_chars.add(s[i])
                if current_substring_length > max_substring_length:
                    max_substring_length = current_substring_length
            i += 1
        return max_substring_length
