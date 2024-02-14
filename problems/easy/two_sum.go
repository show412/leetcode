/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: your name
 * @LastEditTime: 2024-02-14 11:17:32
 * @Description: file content
 */
// https://leetcode.com/problems/two-sum/
// once pass hash
// hasp key store value of nums, hash value store index of nums
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
