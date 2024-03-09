/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: your name
 * @LastEditTime: 2024-03-09 22:45:47
 * @Description: file content
 */
// https://leetcode.com/problems/product-of-array-except-self/
/*
Given an array nums of n integers where n > 1,
return an array output such that output[i]
is equal to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]
Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity?
(The output array does not count as extra space
for the purpose of space complexity analysis.)
*/
/*
前缀乘积和后缀乘积 除去当前这个数
TC is O(n), SC is O(n)
*/
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		product := 1
		if i != 0 {
			product = prefix[i-1] * nums[i-1]
		}
		prefix[i] = product
	}
	for i := len(nums) - 1; i >= 0; i-- {
		product := 1
		if i != len(nums)-1 {
			product = suffix[i+1] * nums[i+1]
		}
		suffix[i] = product
	}
	for i := 0; i < len(nums); i++ {
		res[i] = prefix[i] * suffix[i]
	}

	return res
}

// TC is O(n), SC is O(1)
// 用结果的数组res 存储前缀成绩 然后用一个变量post来记后缀的乘积总和（不算当前的数i）
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		product := 1
		if i != 0 {
			product = res[i-1] * nums[i-1]
		}
		res[i] = product
	}
	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * post
		post = post * nums[i]
	}

	return res
}
