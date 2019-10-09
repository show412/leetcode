// https://leetcode.com/problems/reverse-integer/
/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers
within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem,
assume that your function returns 0 when the reversed integer overflows.
*/
import "math"

func reverse(x int) int {
	res := 0
	for x != 0 {
		units := x % 10
		x = x / 10
		if res > math.MaxInt32/10 || res < math.MinInt32/10 || (res == math.MaxInt32/10 && units > 7) || (res == math.MinInt32/10 && units < -8) {
			return 0
		}
		res = res*10 + units
	}
	return res
}
