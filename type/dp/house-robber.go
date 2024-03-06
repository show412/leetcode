/*
 * @Author: hongwei.sun
 * @Date: 2024-03-06 21:43:17
 * @LastEditors: your name
 * @LastEditTime: 2024-03-06 21:43:18
 * @Description: file content
 */
/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-03-06 21:42:35
 * @Description: file content
 */
// https://leetcode.com/problems/house-robber/
/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them
is that adjacent houses have security system connected
and it will automatically contact the police
if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	f := make([]int, len(nums))
	f[0] = nums[0]
	f[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f[i] = max(f[i-1], f[i-2]+nums[i])
	}
	return f[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
