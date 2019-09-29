package main

import (
	"fmt"
)

/*
test case:
101, 56, 69, 48, 30
4

[8,85,24,85,69]
4

[73,106,39,6,26,15,30,100,71,35,46,112,6,60,110]
29

[98,60,24,89,84,51,61,96,108,87,68,29,14,11,13,50,13,104,57,8,57,111,92,87,9,59,65,116,56,39,55,11,21,105,57,36,48,93,20,94,35,68,64,41,37,11,50,47,8,9]
439
*/
func main() {
	res := minAddToMakeValid("()))((")
	// 998001
	fmt.Println(res)
}

func minAddToMakeValid(S string) int {
	stack := make([]byte, 0)
	for i := len(S) - 1; i >= 0; i-- {
		s := S[i]
		if len(stack) == 0 {
			stack = append(stack, s)
		} else {
			b := stack[len(stack)-1]
			if s == '(' && b == ')' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s)
			}
		}
	}
	return len(stack)
}
