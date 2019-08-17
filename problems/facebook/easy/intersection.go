// https://leetcode.com/problems/intersection-of-two-arrays/
/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Note:

Each element in the result must be unique.
The result can be in any order.
*/
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, 0)
	res := make([]int, 0)
	for _, v := range nums1 {
		m[v] = 1
	}
	for _, v := range nums2 {
		if m[v] == 1 {
			m[v] = 2
		}
	}
	for k, v := range m {
		if v == 2 {
			res = append(res, k)
		}
	}
	return res
}
