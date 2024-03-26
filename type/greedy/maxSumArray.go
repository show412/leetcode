/*
 * @Author: hongwei.sun
 * @Date: 2024-03-02 23:12:05
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-26 15:22:59
 * @Description: file content
 */
// https://leetcode.com/problems/maximum-subarray/
/*
Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach,
which is more subtle.
可以理解为是贪心(动态规划?)算法,
max 的意思是包含第 i 个 num 时的最大值 当max小于0的时候，max只为 num 时max 最大
这就是贪心算法的核心，找子序列sum最大的时候，看看包括每一个值的sum有没有最大的
否则就是当前的 max 加上 num
注意是从第1个开始遍历
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if max <= 0 {
			max = num
		} else {
			max += num
		}
		if max > maxSum {
			maxSum = max
		}
	}
	return maxSum
}