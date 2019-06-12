/*
Given an array of integers and an integer k,
you need to find the total number of continuous subarrays
whose sum equals to k.
Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
*/
// https://leetcode.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			res++
		}
		if i == len(nums)-1 {
			break
		}
		target := k - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target {
				res++
			}
			target -= nums[j]
		}
	}
	return res
}
