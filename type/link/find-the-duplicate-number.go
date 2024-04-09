/*
 * @Author: hongwei.sun
 * @Date: 2024-04-09 23:46:25
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-09 23:52:40
 * @Description: file content
 */
//  https://leetcode.com/problems/find-the-duplicate-number/description/
/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.

 

Example 1:

Input: nums = [1,3,4,2,2]
Output: 2
Example 2:

Input: nums = [3,1,3,4,2]
Output: 3
Example 3:

Input: nums = [3,3,3,3,3]
Output: 3
 

Constraints:

1 <= n <= 105
nums.length == n + 1
1 <= nums[i] <= n
All the integers in nums appear only once except for precisely one integer which appears two or more times.
 

Follow up:

How can we prove that at least one duplicate number must exist in nums?
Can you solve the problem in linear runtime complexity?
*/
/* 
要求不能改nums 并且SC is O(1)
这题有多种解法：
1, 环装链表
将这些数字想象成链表中的结点，数组中数字代表下⼀个结点的数组下
标。找重复的数字就是找链表中成环的那个点。由于题⽬保证了⼀定会有重复的数字，所以⼀定会
成环。所以⽤快慢指针的⽅法，快指针⼀次⾛ 2 步，慢指针⼀次⾛ 1 步，相交以后，快指针从头
开始，每次⾛⼀步，再次遇⻅的时候就是成环的交点处，也即是重复数字所在的地⽅。
2， binary search
	1）. 假设有 n+1 个数，则可能重复的数位于区间 [1, n] 中。记该区间最⼩值、最⼤值和中间值为
low、high、mid
	2）. 遍历整个数组，统计⼩于等于 mid 的整数的个数，⾄多为 mid 个
	3）. 如果超过 mid 个就说明重复的数存在于区间 [low,mid] （闭区间）中；否则，重复的数存在
于区间 (mid, high] （左开右闭）中
	4）. 缩⼩区间，继续重复步骤 2、3，直到区间变成 1 个整数，即 low == high
	5）. 整数 low 就是要找的重复的数
3， 排序， 但是这样的话就把nums改了
*/
// link
func findDuplicate(nums []int) int {

}

// binary search
func findDuplicate(nums []int) int {
	
}

// sort
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return nums[i]
		}
	}
	return len(nums)
}