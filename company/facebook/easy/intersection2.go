// https://leetcode.com/problems/intersection-of-two-arrays-ii/
/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Note:

Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk,
and the memory is limited such that you cannot load all elements into the memory at once?
*/
func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m1 := make(map[int]int, 0)
	m2 := make(map[int]int, 0)
	for _, v := range nums1 {
		if _, ok := m1[v]; !ok {
			m1[v] = 1
		} else {
			m1[v]++
		}
	}
	for _, v := range nums2 {
		if _, ok := m2[v]; !ok {
			m2[v] = 1
		} else {
			m2[v]++
		}
	}
	for k, _ := range m1 {
		if m2[k] >= 1 {
			i := 0
			for i < min(m1[k], m2[k]) {
				res = append(res, k)
				i++
			}
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
