/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-19 10:52:31
 * @Description: file content
 */
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
/*
Given an array of integers nums sorted in ascending order, 
find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/
// it is a common solution to binary search to find range of target
// two binary search to find the boundary left and right
/*
二分查找，尽量往左(不断调整r指针)就能找到第一个元素。
尽量往左(不断调整l指针)就能找到最后一个元素. 
注意在search的时候 mid == target的时候不要退出 继续尽量(往左或者往右)找
*/

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	// find start
	l := 0
	r := len(nums) - 1
	start := -1
	end := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			start = mid
			r = mid - 1
		}
	}
	res[0] = start
	// find end 
	l = 0
	r = len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			end = mid
			l = mid + 1
		}
	}
	res[1] = end
	return res
 }