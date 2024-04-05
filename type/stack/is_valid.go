/*
 * @Author: hongwei.sun
 * @Date: 2024-02-14 15:59:07
 * @LastEditors: your name
 * @LastEditTime: 2024-02-14 16:33:28
 * @Description: file content
 */
/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
*/
// https://leetcode.com/problems/valid-parentheses/
// it's one stack solution
// Initialize an empty stack.
// key point:
// if there is comming one close bracket, the top of stack must be correponding open bracket
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	closeToOpen := map[string]string{")": "(", "}": "{", "]": "["}
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		// notice string to sub array return byte, we need to convert to string
		if openBracket, ok := closeToOpen[string(s[i])]; ok {
			if len(stack) == 0 || openBracket != stack[len(stack)-1] {
				return false
			} else {
				// stack pop last one element
				stack = stack[:len(stack)-1]
			}
		} else {
			// necessary to store the result of append, often in the variable holding the slice itself
			stack = append(stack, string(s[i]))
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
