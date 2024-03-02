/*
 * @Author: hongwei.sun
 * @Date: 2024-03-02 11:42:02
 * @LastEditors: your name
 * @LastEditTime: 2024-03-02 13:36:59
 * @Description: file content
 */
//  https://leetcode.com/problems/min-cost-climbing-stairs/description/
//  You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

//  You can either start from the step with index 0, or the step with index 1.

//  Return the minimum cost to reach the top of the floor.

//  Example 1:

//  Input: cost = [10,15,20]
//  Output: 15
//  Explanation: You will start at index 1.
//  - Pay 15 and climb two steps to reach the top.
//  The total cost is 15.
//  Example 2:

//  Input: cost = [1,100,1,1,1,100,1,1,100,1]
//  Output: 6
//  Explanation: You will start at index 0.
//  - Pay 1 and climb two steps to reach index 2.
//  - Pay 1 and climb two steps to reach index 4.
//  - Pay 1 and climb two steps to reach index 6.
//  - Pay 1 and climb one step to reach index 7.
//  - Pay 1 and climb two steps to reach index 9.
//  - Pay 1 and climb one step to reach the top.
//  The total cost is 6.

// Constraints:

// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999
func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	// f[n] is min cost to arrive n from bottom
	f := make([]int, l)
	f[0] = cost[0]
	f[1] = cost[1]

	for i := 2; i < l; i++ {
		f[i] = min(f[i-1]+cost[i], f[i-2]+cost[i])
	}
	return min(f[l-1], f[l-2])
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}