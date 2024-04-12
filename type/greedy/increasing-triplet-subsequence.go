/*
 * @Author: hongwei.sun
 * @Date: 2024-04-12 22:21:55
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-12 22:26:57
 * @Description: file content
 */
// https://leetcode.com/problems/increasing-triplet-subsequence/description/
/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

 

Example 1:

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
 

Constraints:

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1
*/
/*
也算贪心的一种
找到small 和 mid 再找到比mid还大的 那就肯定存在一个三元组
*/
func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
		return false
	}
	small := math.MaxInt32
	mid := math.MaxInt32
	/*
	if else 也是有先后顺序的，如果存在三元组，那肯定是先找到最小的, 再找到中间的
	最后存在一个都比small和mid大的数
	*/
	for _, num := range nums {
		if num <= small {
			small = num
		} else if num <= mid {
			mid = num
		} else {
			return true
		}
	}
	return false
}