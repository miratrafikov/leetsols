class Solution:
    def validate_one_way(self, s: str, t: str) -> bool:
        chars_last_occurence_in_s = {}
        for i in range(len(s)):
            if s[i] not in chars_last_occurence_in_s:
                chars_last_occurence_in_s[s[i]] = i
            else:
                if t[chars_last_occurence_in_s[s[i]]] != t[i]:
                    return False
        return True
    
    def isIsomorphic(self, s: str, t: str) -> bool:
        return self.validate_one_way(s, t) and self.validate_one_way(t, s)
