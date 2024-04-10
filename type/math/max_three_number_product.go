/*
 * @Author: hongwei.sun
 * @Date: 2024-04-10 22:27:09
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 22:28:42
 * @Description: file content
 */

//  https://leetcode.com/problems/maximum-product-of-three-numbers/
/*
Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

 

Example 1:

Input: nums = [1,2,3]
Output: 6
Example 2:

Input: nums = [1,2,3,4]
Output: 24
Example 3:

Input: nums = [-1,-2,-3]
Output: -6
 

Constraints:

3 <= nums.length <= 104
-1000 <= nums[i] <= 1000
*/
/*
这个的问题是要考虑负数的情况
所以最后要比较两个最小的数和最大数乘积 以及三个最大数的乘积
*/
func maximumProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	if len(nums) <= 3 {
		for i := 0; i < len(nums); i++ {
			res = res * nums[i]
		}
		return res
	}
	sort.Ints(nums)
	return max(nums[0]*nums[1]*nums[len(nums)-1],
		nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}
 func max(a, b int) int {
	if a > b {
		return a
	}
	return b
 }