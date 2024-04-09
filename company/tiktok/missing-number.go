/*
 * @Author: hongwei.sun
 * @Date: 2024-04-09 15:04:36
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-09 15:25:36
 * @Description: file content
 */
// https://leetcode.com/problems/missing-number/
/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, 
find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/
func missingNumber(nums []int) int {
	sort.Ints(nums)
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == n {
			n++
			continue
		} else {
			return n
		}
	}
	// if no any miss, means miss last one len(nums)
	return len(nums)
}
