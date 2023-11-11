# Solution 2: don't start all over, instead start from the previous occurence of the current char
# (and do reset the substring's map and length)
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_substring_length = 0
        current_substring_length = 0
        current_substring_char_to_position = {}
        i = 0
        while i < len(s):
            if s[i] in current_substring_char_to_position:
                i = current_substring_char_to_position[s[i]] + 1
                current_substring_char_to_position = {}
                current_substring_length = 0
            else:
                current_substring_char_to_position[s[i]] = i
                current_substring_length += 1
                if current_substring_length > max_substring_length:
                    max_substring_length = current_substring_length
                i += 1
        return max_substring_length 
