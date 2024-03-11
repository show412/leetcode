/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-03-11 13:13:30
 * @Description: file content
 */
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/
/*
two pointers solution
l is from beginning, r is from end.
if l+r is greater than target, move r otherwise move l
*/
func twoSum(numbers []int, target int) []int {
	res := make([]int, 0)
	l := 0
	r := len(numbers) - 1
	for l <= r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			res = []int{l + 1, r + 1}
			break
		}
	}
	return res
}