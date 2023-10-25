class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        m = {}
        for letter in s:
            if letter in m:
                m[letter] += 1
            else:
                m[letter] = 1
        for letter in t:
            if letter in m:
                if m[letter] == 1:
                    del m[letter]
                else:
                    m[letter] -= 1
            else:
                return False
        return len(m) == 0
