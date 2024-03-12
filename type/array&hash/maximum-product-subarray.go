/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: your name
 * @LastEditTime: 2024-03-12 10:49:12
 * @Description: file content
 */
// https://leetcode.com/problems/maximum-product-subarray/
/*
Given an integer array nums, find the contiguous subarray
within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/
// 也是一种贪心的思想 和maxSumArray不同 那里判断的是max是否大于0 这里是当前遍历到的元素是否大于0
// 这里要记两个变量 最大和最小乘积
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res, imax, imin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			// when num is smaller than 0, max multiply it will be more smaller
			// so swap max and min
			imax, imin = imin, imax
		}
		imax = max(num, num*imax)
		imin = min(num, num*imin)
		res = max(res, imax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//  this solution is wrong, can't work
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	imax, imin := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			imax = max(imax, num*imin, num)
			imin = min(num*imax, imin, num)
		} else {
			imax = max(imax, num*imax, num)
			imin = min(num*imin, imin, num)
		}
	}
	return imax
}
func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}