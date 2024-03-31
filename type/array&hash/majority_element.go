/*
 * @Author: hongwei.sun
 * @Date: 2024-03-31 20:46:58
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-31 20:46:58
 * @Description: file content
 */
/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-31 20:36:42
 * @Description: file content
 */
// https://leetcode.com/problems/majority-element/description/
// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

 

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
 

// Constraints:

// n == nums.length
// 1 <= n <= 5 * 104
// -109 <= nums[i] <= 109

func majorityElement(nums []int) int {
	return nil if nums == nil || len(nums) == 0
	var candidate int
	var count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}