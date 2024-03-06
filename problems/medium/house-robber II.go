/*
 * @Author: hongwei.sun
 * @Date: 2024-03-06 21:46:05
 * @LastEditors: your name
 * @LastEditTime: 2024-03-06 22:20:27
 * @Description: file content
 */
/*
https://leetcode.com/problems/house-robber-ii/
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed.
All houses at this place are arranged in a circle.
That means the first house is the neighbor of the last one. Meanwhile,
adjacent houses have a security system connected,
and it will automatically contact the police
if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.


Example 1:

Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
because they are adjacent houses.
Example 2:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 3:

Input: nums = [1,2,3]
Output: 3


Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 1000

*/
/*
we can reuse house-robber I solution
actually there are two scenarios: including nums[0] and not including nums[0]
if including nums[0], we need to remove nums[len(nums)-1]
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob_helper(nums[1:]), rob_helper(nums[:len(nums)-1]))
}
func rob_helper(nums []int) int {
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