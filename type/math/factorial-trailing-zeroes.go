/*
 * @Author: hongwei.sun
 * @Date: 2024-04-13 09:34:47
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-13 09:42:46
 * @Description: file content
 */
//  https://leetcode.com/problems/factorial-trailing-zeroes/description/
 /*
 Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

 

Example 1:

Input: n = 3
Output: 0
Explanation: 3! = 6, no trailing zero.
Example 2:

Input: n = 5
Output: 1
Explanation: 5! = 120, one trailing zero.
Example 3:

Input: n = 0
Output: 0
 

Constraints:

0 <= n <= 104
 

Follow up: Could you write a solution that works in logarithmic time complexity?
 */
 /*
 有多少个0 取决于有多少个2和5中的最小值
 */
 func trailingZeroes(n int) int {
	divide2 := 0
	divide5 := 0
    num2 := n
    num5 := n
	for num2 != 0 {
		divide2 += num2/2
		num2 = num2/2
	}
	for num5 != 0 {
		divide5 += num5/5
		num5 = num5/5
	}
	return min(divide2, divide5)
}