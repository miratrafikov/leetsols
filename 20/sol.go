type stack []rune

func (s *stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *stack) Pop() rune {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func isValid(s string) bool {
	stack := stack{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.Push(c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last_opening := stack.Pop()
			if !(c == ')' && last_opening == '(' || c == '}' && last_opening == '{' || c == ']' && last_opening == '[') {
				return false
			}
		}
	}
	return len(stack) == 0
}
