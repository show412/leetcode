/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
*/
// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	bracket := map[string]string{"(": ")", "{": "}", "[": "]"}
	stack := make([]string, 0)
	for i := len(s) - 1; i >= 0; i-- {
		sub := string(s[i])
		if v, ok := bracket[sub]; ok {
			if len(stack) == 0 {
				return false
			}
			close := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if v != close {
				return false
			}
		} else {
			stack = append(stack, sub)
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
