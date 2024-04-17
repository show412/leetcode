/*
 * @Author: hongwei.sun
 * @Date: 2024-04-17 14:52:14
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-17 15:37:00
 * @Description: file content
 */
//  https://leetcode.com/problems/sqrtx/
/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
 

Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
 

Constraints:

0 <= x <= 231 - 1
*/
/*
如果x>=4 sqrt x < x/2 
所以通过 1到x/2去二分
*/
func mySqrt(x int) int {
    if x == 0 {
		return 0
	}
	// 1 - 3 都是 1
	if x < 4 {
		return 1
	}
	l := 1
	r := x/2
    res := 1
	for l <= r {
		mid := l + (r-l)/2
		if mid < x/mid {
			l = mid + 1
            res = mid
		} else if mid > x/mid {
			r = mid - 1
		} else {
			return mid
		}
	}
	return res
}