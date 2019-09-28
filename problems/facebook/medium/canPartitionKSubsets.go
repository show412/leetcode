import "sort"

// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
/*
Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.



Example 1:

Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
Output: True
Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.


Note:

1 <= k <= len(nums) <= 16.
0 < nums[i] < 10000.
*/
/*
不能用dp了 用dfs解出
refer to https://www.cnblogs.com/grandyang/p/7733098.html
*/
func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	return dfs(nums, k, sum/k, 0, 0, visited)
}

func dfs(nums []int, k int, target int, start int, curSum int, visited []bool) bool {
	if k == 1 {
		return true
	}
	if curSum > target {
		return false
	}
	if curSum == target {
		return dfs(nums, k-1, target, 0, 0, visited)
	}
	for i := start; i < len(nums); i++ {
		if visited[i] == true {
			continue
		}
		visited[i] = true
		if dfs(nums, k, target, i+1, curSum+nums[i], visited) {
			return true
		}
		visited[i] = false
	}
	return false
}


