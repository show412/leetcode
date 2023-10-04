// https://leetcode.com/problems/subarray-sum-equals-k/
/*
Given an array of integers and an integer k,
you need to find the total number of continuous subarrays
whose sum equals to k.
Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
*/
// TC O(n^2) SC O(1)
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

//  use hashmap TC O(n), SC O(n)
//  depends on the math: a[i] - a[j] = k means the sum is k from index j to index i
// test case [0,0,0,0,0,0,0,0,0,0]  0   result:55
func subarraySum(nums []int, k int) int {
	res := 0
	// key is sum, value is count, the defination is critica
	sumMap := make(map[int]int, 0)
	sumMap[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		sum += num
		if _, ok := sumMap[sum-k]; ok {
			// res should plus the sumMap[sum-k]
			res += sumMap[sum-k]
		}
		sumMap[sum] += 1
	}
	return res
}