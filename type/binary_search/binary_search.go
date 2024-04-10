/*
 * @Author: hongwei.sun
 * @Date: 2024-04-10 12:18:20
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 12:18:20
 * @Description: file content
 */
//  Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

//  You must write an algorithm with O(log n) runtime complexity.

//  Example 1:

//  Input: nums = [-1,0,3,5,9,12], target = 9
//  Output: 4
//  Explanation: 9 exists in nums and its index is 4
//  Example 2:

//  Input: nums = [-1,0,3,5,9,12], target = 2
//  Output: -1
//  Explanation: 2 does not exist in nums so return -1

//  Constraints:

//  1 <= nums.length <= 104
//  -104 < nums[i], target < 104
//  All the integers in nums are unique.
//  nums is sorted in ascending order.

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	//1， if only one element in nums, left == right
	for left <= right {
		//2， to avoid overflow, we can't use (left + right)/2
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 3, that means all elements before mid are not target, remove them, so left = mid + 1
			left = mid + 1
			// 4, that means all elements after mid are not target, remove them, so right = mid -1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// there is hidden tips, finally left and right must be possiblities:
			// [left, mid, right] or [left, right] or [left(right)] or [right, left](out of loop)
			return mid
		}
	}
	return -1
}