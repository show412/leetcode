import "unicode"

// https://leetcode.com/problems/basic-calculator-ii/
/*
Implement a basic calculator to evaluate a simple expression string.

The expression string contains only non-negative integers,
+, -, *, / operators and empty spaces .
The integer division should truncate toward zero.

Example 1:

Input: "3+2*2"
Output: 7
Example 2:

Input: " 3/2 "
Output: 1
Example 3:

Input: " 3+5 / 2 "
Output: 5
Note:

You may assume that the given expression is always valid.
Do not use the eval built-in library function.
*/
func calculate(s string) int {
	l1, o1, l2, o2 := 0, 1, 1, 1
	if s[0] == '-' {
		s = "0" + s
	}
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if unicode.IsDigit(c) {
			num := int(c - '0')
			for (i+1) < len(s) && unicode.IsDigit(rune(s[i+1])) {
				num = num*10 + int(s[i+1]-'0')
				i++
			}
			if o2 == 1 {
				l2 = l2 * num
			} else {
				l2 = l2 / num
			}
		} else if c == '*' || c == '/' {
			if c == '*' {
				o2 = 1
			} else {
				o2 = -1
			}
		} else if c == '+' || c == '-' {
			l1 = l1 + o1*l2
			if c == '+' {
				o1 = 1
			} else {
				o1 = -1
			}
			o2 = 1
			l2 = 1
		}
	}
	return l1 + o1*l2
}
