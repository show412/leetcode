/*
 * @Author: hongwei.sun
 * @Date: 2024-02-13 16:58:48
 * @LastEditors: your name
 * @LastEditTime: 2024-02-13 17:02:31
 * @Description: file content
 Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, 0)
	ret := false
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == true {
			ret = true
			break
		} else {
			m[nums[i]] = true
		}
	}
	return ret
}