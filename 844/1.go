type stack []rune

func (s *stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *stack) Pop() rune {
	n := len(*s)
	v := (*s)[n - 1]
	*s = (*s)[:n - 1]
	return v
}

func backspaceCompare(s string, t string) bool {
    var sWithBackspacesApplied stack
    var tWithBackspacesApplied stack
    for _, char := range s {
        if char != '#' {
            sWithBackspacesApplied.Push(char)
        } else if len(sWithBackspacesApplied) != 0 {
            sWithBackspacesApplied.Pop()
        }
    }
    for _, char := range t {
        if char != '#' {
            tWithBackspacesApplied.Push(char)
        } else if len(tWithBackspacesApplied) != 0 {
            tWithBackspacesApplied.Pop()
        }
    }
    if len(sWithBackspacesApplied) != len(tWithBackspacesApplied) {
        return false
    }
    for i := 0; i < len(sWithBackspacesApplied); i++ {
        if sWithBackspacesApplied[i] != tWithBackspacesApplied[i] {
            return false
        }
    }
    return true
}
