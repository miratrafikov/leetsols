package main

import "fmt"

func main() {
	testcases := [][]string{
		[]string{"abc", "abc"},
		[]string{"a_bc", "a_bc"},
		[]string{"a__bc", "a_bc"},
		[]string{"a__bc__", "a_bc_"},
		[]string{"_", "_"},
		[]string{"_____", "_"},
	}

	for i := 0; i < len(testcases); i++ {
		isCorrect := testcases[i][1] == problem(testcases[i][0])
		fmt.Println(problem(testcases[i][0]))
		fmt.Println(isCorrect)
	}
}

func problem(s string) string {
	n := len(s)
	l, r := 0, 0
	removals := 0
	arr := make([]byte, len(s))
	for l < n {
		arr[l-removals] = s[l]
		if s[l] != '_' {
			l++
			r++
		} else {
			r++
			for r < n && s[r] == '_' {
				r++
				removals++
			}
			l = r
		}
	}
	return string(arr[:n-removals])
}
