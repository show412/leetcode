/*
 * @Author: hongwei.sun
 * @Date: 2024-03-06 21:32:10
 * @LastEditors: your name
 * @LastEditTime: 2024-03-06 21:32:11
 * @Description: file content
 */
/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: your name
 * @LastEditTime: 2024-02-29 23:15:59
 * @Description: file content
 */
// https://leetcode.com/problems/climbing-stairs/
/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	f := make([]int, 0)
	f[0] = 1
	f[1] = 1
	// https://www.spiceworks.com/tech/devops/articles/what-is-dynamic-programming/
	// DP solution i solution is i-1 solution (1 step again will be i) + i-2 solution (2 steps again will be i)
	for i := 2; i < n+1; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}