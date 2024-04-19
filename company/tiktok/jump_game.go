/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-31 11:35:49
 * @Description: file content
 */
// https://leetcode.com/problems/jump-game/
/*
Given an array of non-negative integers, 
you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
*/

// greedy algo solution
func canJump(nums []int) bool {
	// goal is what we want to reach
	// backfowards to begining
	goal := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		// that means i position + what steps it can jump greater or equal than goal
		// this i will be next goal
		if nums[i]+i >= goal {
			goal = i
		}
	}
	if goal == 0 {
		return true
	}
	return false
}

// DP?
func canJump(nums []int) bool {
	f := make([]bool, len(nums))
	f[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// can jump to i from j
			f[i] = f[j] && nums[j] >= (i-j)
			if f[i] == true {
				break
			}
		}
	}
	return f[len(nums)-1]
}