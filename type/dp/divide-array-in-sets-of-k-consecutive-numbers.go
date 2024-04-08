/*
 * @Author: hongwei.sun
 * @Date: 2024-04-08 23:03:31
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-08 23:04:39
 * @Description: file content
 */
// https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/description/
/*
Given an array of integers nums and a positive integer k, check whether it is possible to divide this array into sets of k consecutive numbers.

Return true if it is possible. Otherwise, return false.

 

Example 1:

Input: nums = [1,2,3,3,4,4,5,6], k = 4
Output: true
Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].
Example 2:

Input: nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
Output: true
Explanation: Array can be divided into [1,2,3] , [2,3,4] , [3,4,5] and [9,10,11].
Example 3:

Input: nums = [1,2,3,4], k = 3
Output: false
Explanation: Each array should be divided in subarrays of size 3.
 

Constraints:

1 <= k <= nums.length <= 105
1 <= nums[i] <= 109
*/
func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	m := make(map[int]int)
	sort.Ints(nums)
	for _, num := range nums {
		m[num]++
	}
	for i := 0; i<len(nums);i++ {
		if m[nums[i]] > 0 {
			for j := 0; j < k; j++ {
				if m[nums[i]+j] <= 0 {
					return false
				} else {
					m[nums[i]+j]--
				}
			}
		}
	}
	return true
}