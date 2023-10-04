// https://leetcode.com/problems/rotate-array
/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
Note:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?
*/
func rotate(nums []int, k int) {
	if k == 0 || k == len(nums) {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	// The nums is slice, slice is a pointer which point to one array,
	// so it doesn't need to return nums.
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	i := start
	j := end
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
