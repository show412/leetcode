/*
 * @Author: hongwei.sun
 * @Date: 2024-03-26 10:36:48
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-26 10:51:44
 * @Description: file content
 */
//  https://leetcode.com/problems/target-sum/description/
// You are given an integer array nums and an integer target.

// You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

// For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
// Return the number of different expressions that you can build, which evaluates to target.

 

// Example 1:

// Input: nums = [1,1,1,1,1], target = 3
// Output: 5
// Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
// Example 2:

// Input: nums = [1], target = 1
// Output: 1
 

// Constraints:

// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000
/*
typical dfs solution
from 0 index, two choice + or -
we just brute force it
*/
var res int
func findTargetSumWays(nums []int, target int) int {
    res = 0
	dpf(0,0,nums,target)
	return res
}

func dpf(i int, sum int, nums []int, target int) {
	if i == len(nums) {
		if sum == target {
			res++
		}
		return
	}
	dpf(i+1, sum+nums[i],nums, target)
	dpf(i+1, sum-nums[i],nums, target)
	return
}